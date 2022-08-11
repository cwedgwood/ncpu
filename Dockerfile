FROM scratch
COPY ./n /n
ENTRYPOINT [ "./n" ]
