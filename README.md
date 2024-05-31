# zap-yaegi-issue 

Just a demo to illustrate issues experienced with zap when using inside a traffic plugin.

## Repro

1. Start whoami container which serves as backend service `podman run --name whoami -d --rm  -p 8080:8888 -p 8081:80 traefik/whoami`
2. Start traefik container with plugin: `podman run -it --rm -v $PWD/traefik.yaml:/etc/traefik.yaml -v $PWD/src:/plugins-local/src --network container:whoami docker.io/traefik:v2.11.0 --configfile=/etc/traefik.yaml`
3. In another terminal window `curl http://localhost:8080/demo`