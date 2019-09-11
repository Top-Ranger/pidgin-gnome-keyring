# Original version by Ali Ebrahim <ali.ebrahim314@gmail.com>
# Modifications:
#     2018 Marcus Soll: Changed file for kwallet usage
VERSION = 0.1.0
PURPLEFLAGS = `pkg-config --cflags purple`

all: kwallet-dbus-interface.a pidgin-kwallet.so

clean:
	rm -f pidgin-kwallet.so kwallet-dbus-interface.h kwallet-dbus-interface.a

kwallet-dbus-interface.a:
	go build -buildmode=c-archive kwallet-dbus-interface/kwallet-dbus-interface.go

pidgin-kwallet.so:
	gcc -pthread -Wall -O2 -shared -fPIC -DPIC pidgin-kwallet.c kwallet-dbus-interface.a -o pidgin-kwallet.so ${PURPLEFLAGS} -DVERSION=\"${VERSION}\"

install: pidgin-kwallet.so
	mkdir -p ${DESTDIR}/usr/lib/purple-2/
	cp pidgin-kwallet.so ${DESTDIR}/usr/lib/purple-2/

install_local: pidgin-kwallet.so
	mkdir -p ~/.purple/plugins
	cp pidgin-kwallet.so ~/.purple/plugins/
