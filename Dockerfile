#FROM golang:1.16 AS builder

#EXPOSE 1551
#EXPOSE 8080



#COPY assets/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

#COPY . /app
#RUN make /app/
#RUN go mod tidy
#RUN go get github.com/suchy1105/GUIcontroler &&  go get github.com/go-chi/chi && go mod init github.com/suchy1105/GUIcontroler && go build -o /app/dist/ -v /app/cmd/gui/main.go
#RUN chmod -R +x /opt/h/* && mkdir -p /var/husar/gui

#VOLUME /home/karoldb/GolangProjects/GUIsocket
FROM golang:1.17
#AS executor
#COPY --from=builder /app /app

# install packages and set Warsaw timezone
RUN printf '%s\n' 'path-exclude /usr/share/doc/*' 'path-include /usr/share/doc/*/copyright' 'path-exclude /usr/share/man/*' 'path-exclude /usr/share/groff/*' 'path-exclude /usr/share/info/*' 'path-exclude /usr/share/lintian/*' 'path-exclude /usr/share/linda/*' > /etc/dpkg/dpkg.cfg.d/01_nodoc && \
    echo 'APT::Install-Recommends "0" ; APT::Install-Suggests "0" ;' >> /etc/apt/apt.conf && export DEBIAN_FRONTEND=noninteractive && \
    sed -i -e 's/main/main contrib/g' /etc/apt/sources.list && \
    apt-get update && apt-get install --no-install-recommends -yq firefox-esr x11-apps xauth tzdata apt-utils mesa-utils and libgl1-mesa-glx wayland-protocols libxkbcommon-x11-dev libwayland-egl1 libwayland-dev libwayland-bin libwayland-egl1-mesa libfontconfig1 ca-certificates fontconfig ttf-mscorefonts-installer imagemagick x11-xserver-utils apt-utils libegl1 libegl1-mesa libegl-mesa0 make golang && \
    fc-cache  && cp /usr/share/zoneinfo/Europe/Warsaw /etc/localtime && \
    echo "Europe/Warsaw" >  /etc/timezone
#RUN apt-get update && apt-get install  gettext-base git-daemon-run git
#| git-daemon-sysvinit git-doc git-el git-email
#RUN apt-get update && apt-get install git-gui gitk gitweb git-cvs git-mediawiki git-svn perl-doc
#RUN apt-get update && apt-get install libterm-readline-gnu-perl | libterm-readline-perl-perl libb-debug-perl
#RUN apt-get update && apt-get install liblocale-codes-perl

COPY . /app
WORKDIR /app
RUN  go get github.com/go-chi/chi
#&& go mod tidy

RUN go mod download
#github.com/suchy1105/GUIcontroler
#RUN go mod init github.com/suchy1105/GUIcontroler

RUN go build -o /app/dist/ -v /app/cmd/gui/main.go
RUN chmod -R +x /opt/h/* && mkdir -p /var/husar/gui

STOPSIGNAL SIGINT
CMD ["/app/dist/dist"]
