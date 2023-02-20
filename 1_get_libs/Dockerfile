FROM golang:latest as build


FROM gcr.io/distroless/base

#bash
COPY --from=build /bin/bash /bin/bash
COPY --from=build /lib/x86_64-linux-gnu/libdl.so.2 /lib/x86_64-linux-gnu/libdl.so.2
COPY --from=build /usr/lib/x86_64-linux-gnu/libdl.so /usr/lib/x86_64-linux-gnu/libdl.so
COPY --from=build /lib/x86_64-linux-gnu/libtinfo.so.6 /lib/x86_64-linux-gnu/libtinfo.so.6
COPY --from=build /lib/x86_64-linux-gnu/libtinfo.so.6.2 /lib/x86_64-linux-gnu/libtinfo.so.6.2

#sh
COPY --from=build /bin/sh /bin/sh

#tar
COPY --from=build /bin/tar /bin/tar
COPY --from=build /usr/include/tar.h /usr/include/tar.h
COPY --from=build /usr/share/man/man1/tar.1.gz /usr/share/man/man1/tar.1.gz
#for tar
COPY --from=build /usr/lib/x86_64-linux-gnu/libacl.so.1 /usr/lib/x86_64-linux-gnu/libacl.so.1
COPY --from=build /usr/lib/x86_64-linux-gnu/libacl.so.1.1.2253 /usr/lib/x86_64-linux-gnu/libacl.so.1.1.2253
COPY --from=build /lib/x86_64-linux-gnu/libselinux.so.1 /lib/x86_64-linux-gnu/libselinux.so.1
COPY --from=build /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/libc.so.6
COPY --from=build /lib/x86_64-linux-gnu/libc-2.31.so /lib/x86_64-linux-gnu/libc-2.31.so
COPY --from=build /usr/lib/x86_64-linux-gnu/libpcre2-8.so.0 /usr/lib/x86_64-linux-gnu/libpcre2-8.so.0
COPY --from=build /usr/lib/x86_64-linux-gnu/libpcre2-8.so.0.10.1 /usr/lib/x86_64-linux-gnu/libpcre2-8.so.0.10.1
COPY --from=build /lib/x86_64-linux-gnu/libdl.so.2 /lib/x86_64-linux-gnu/libdl.so.2
COPY --from=build /lib/x86_64-linux-gnu/libdl-2.31.so /lib/x86_64-linux-gnu/libdl-2.31.so
COPY --from=build /lib64/ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2
COPY --from=build /lib/x86_64-linux-gnu/ld-2.31.so /lib/x86_64-linux-gnu/ld-2.31.so
COPY --from=build /lib/x86_64-linux-gnu/libpthread.so.0 /lib/x86_64-linux-gnu/libpthread.so.0
COPY --from=build /lib/x86_64-linux-gnu/libpthread-2.31.so /lib/x86_64-linux-gnu/libpthread-2.31.so

#gzip
COPY --from=build /lib/x86_64-linux-gnu/libresolv-2.31.so /lib/x86_64-linux-gnu/libresolv-2.31.so
COPY --from=build /bin/gzip /bin/gzip 
COPY --from=build /usr/share/man/man1/gzip.1.gz /usr/share/man/man1/gzip.1.gz
COPY --from=build /usr/share/info/gzip.info.gz /usr/share/info/gzip.info.gz
COPY --from=build /bin/gzip /usr/bin/gzip

#find
COPY --from=build /usr/bin/find /usr/bin/find
COPY --from=build /usr/share/man/man1/find.1.gz /usr/share/man/man1/find.1.gz
COPY --from=build /usr/share/info/find.info-2.gz /usr/share/info/find.info-2.gz
COPY --from=build /usr/share/info/find.info.gz /usr/share/info/find.info.gz
COPY --from=build /usr/share/info/find.info-1.gz /usr/share/info/find.info-1.gz

#mkdir
COPY --from=build /usr/share/man/man1/mkdir.1.gz /usr/share/man/man1/mkdir.1.gz
COPY --from=build /bin/mkdir /bin/mkdir

#ls
COPY --from=build /bin/ls /bin/ls
COPY --from=build /usr/share/man/man1/ls.1.gz /usr/share/man/man1/ls.1.gz

#cp
COPY --from=build /bin/cp /bin/cp
COPY --from=build /usr/share/man/man1/cp.1.gz /usr/share/man/man1/cp.1.gz
COPY --from=build /usr/lib/x86_64-linux-gnu/libattr.so.1 /usr/lib/x86_64-linux-gnu/libattr.so.1
COPY --from=build /usr/lib/x86_64-linux-gnu/libattr.so.1.1.2448 /usr/lib/x86_64-linux-gnu/libattr.so.1.1.2448
