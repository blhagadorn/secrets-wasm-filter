# Secrets WASM Filter
WASM Envoy filter for injecting secrets into requests dynamically.



## Steps to accomplish this:
1. Create test API (Node or Golang) that will enforce authorization requests
2. Create client container with sidecar and attach filter deployment
3. Get POC working with cURL command from client container with secret mounted in Envoy directly.
4. Inject secret with Vault with https://www.vaultproject.io/docs/platform/k8s/injector
5. Harden SYS_CALLs on Envoy attached sidecar

Auth Types to try:

Basic Auth
JWT
Dynamically change secrets without client being aware


### Initial Inspiration
```
Hardened sidecar container (with filter) injects secrets onto requests dynamically Rather than storing tokens in the application container in k8s pod (in ENV or , they could be stored in the filter container and attached to requests as they are made.  For example the application container would make request example.com?token=TOKEN and then the filter would tack on the TOKEN as the request is made without needing to store it in the main container. Hashicorp Vault could potentially bootstrap the proxy container rather than the main application container like is currently recommended. This container could be extremely hardened (Linux capabilities and other items locked down as well as strict Allowed List filters for the resources it needs to hit) to be just good at tacking on secrets to requests. Could tack onto RESTful requests and others.  With quick deployment of the filter, secrets could be dynamically rotated often without the application needing to rotate them as well.
```
