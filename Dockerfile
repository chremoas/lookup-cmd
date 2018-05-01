FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>

ADD lookup-cmd-linux-amd64 lookup-cmd
VOLUME /etc/chremoas

ENTRYPOINT ["/lookup-cmd", "--configuration_file", "/etc/chremoas/auth-bot.yaml"]