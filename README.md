**In development**

****

# Docker Sphere

Docker Sphere is a docker container management solution that provides
an easy way to manage your docker deployments.

## Adding protocol buffers

To add a new protocol buffer, you need to add the `.proto` file to the `protos` directory. Then, you need to run the
following command to generate the GO code from the protocol buffer:

```bash
bash ./scripts/build-protos.sh
```