# Test GO application consist of server and client to test service mesh

## Run with skaffold:
1. git clone https://github.com/MichaelGit-beep/go_app_compote.git
2. Change docker images to use to match your image registry in this files:
- skaffold.yaml
- server/Dockerfile
- server/k8s/pod.yaml
- client/Dockerfile
- client/k8s/pod.yaml