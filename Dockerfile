FROM debian:buster

EXPOSE 1551
EXPOSE 8080

# install packages and set Warsaw timezone
RUN printf '%s\n' 'path-exclude /usr/share/doc/*' 'path-include /usr/share/doc/*/copyright' 'path-exclude /usr/share/man/*' 'path-exclude /usr/share/groff/*' 'path-exclude /usr/share/info/*' 'path-exclude /usr/share/lintian/*' 'path-exclude /usr/share/linda/*' > /etc/dpkg/dpkg.cfg.d/01_nodoc && \
    echo 'APT::Install-Recommends "0" ; APT::Install-Suggests "0" ;' >> /etc/apt/apt.conf && export DEBIAN_FRONTEND=noninteractive && \
    sed -i -e 's/main/main contrib/g' /etc/apt/sources.list && \
    apt-get update && apt-get install --no-install-recommends -yq tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Europe/Warsaw /etc/localtime && \
    echo "Europe/Warsaw" >  /etc/timezone

COPY assets/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

#COPY dist/hdetect /opt/husar/hdetect
#COPY dist/hdetect_launcher /opt/husar/hdetect_launcher

RUN chmod -R +x /opt/husar/* && chmod -R +x /usr/bin

VOLUME /etc/husar
VOLUME /var/husar

STOPSIGNAL SIGINT
CMD ["/opt/husar/hdetect_launcher", "/opt/husar/hdetect"]
