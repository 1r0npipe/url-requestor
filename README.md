# URL requestor

## To run the applicatoin:
<code>
git clone https://github.com/1r0npipe/url-requestor
</code> 

<code>cd url-requestor</code><br>
<code>make run
</code>

## To run tests
under maintenance

## The manifest files are at *k8s* folder
Apply files for available Kubernetes environment (make sure image name is correct, since there is not setup as might change by tag):<br>
<code>
kubectl -f apply url-requestor-deployment.yaml</code><br>
<code>
kubectl -f apply url-requestor-svc.yaml
</code>
## Current status
1. Server is running
2. Config file is correctly reading (in YAML format)
3. Makefile is ready
4. K8s templates for deployment and service is ready
5. Server is available to check the status
6. Simple logger middleware is applied
7. Test event file is ready
8. Dockerfile is ready and working well
9. Defined several typical errors
10. Do the main functionality with all checks and sortings

## TODO:

1. Make unit tests
2. finish README file
