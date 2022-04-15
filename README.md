# URL requestor

## To run the applicatoin:
<code>
git clone https://github.com/1r0npipe/url-requestor
</code> 

<code>cd url-requestor</code><br>
<code>make run
</code>

Pay attention at the config file (./configs/config.yaml), there you can find the array object of requested URLs, you can add or modify that list if you need to request more/less URLs with similar JSON answer

## To run tests inside of project directory
<code>make test
</code>

To verify its work you can run this at the browser:
<code>http://0.0.0.0:8080/request?sortKey=views&limit=2</code>
or run at termital the command:
<code>
curl -XGET 'http://0.0.0.0:8080/request?sortKey=views&limit=2'
</code>
you should see somethin like this:
<code>
{"data":[{"url":"www.example.com/abc1","views":1000,"relevanceScore":0.1},{"url":"www.example.com/abc2","views":2000,"relevanceScore":0.2}]}
</code>
## The manifest files are at *k8s* folder
Apply files for available Kubernetes environment (make sure image name is correct, since there is not setup as might be changed by tag):<br>
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
