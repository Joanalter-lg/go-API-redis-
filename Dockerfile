# Runtime-only image with no external base pull.
FROM scratch

ARG PORT=8080
ENV PORT=${PORT}

COPY main /main

EXPOSE ${PORT}
ENTRYPOINT ["/main"]
