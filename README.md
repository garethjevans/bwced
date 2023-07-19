# bwced 

a very (very) simple storage server.

this should not be used for anything production!

## Installation

To install `bwced` on a local k8s cluster, use: 

```shell
kubectl apply -f resources/server-it.yaml
```

This will create a `bwced` namespace and install the application there.

To run a post deployment tekton task to validate that the server is functioning.

```shell
kubectl create -f resources/test-taskrun.yaml
tkn taskrun logs -f
```

You should see logs that look something like:

```shell
taskrun.tekton.dev/bwced-kcfpz created
? Select taskrun: bwced-kcfpz started 1 second ago
[upload] Generating test content...
[upload] test content
[upload] Uploading file...
[upload]   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
[upload]                                  Dload  Upload   Total   Spent    Left  Speed
100   235  100    36  100   199   6888  38078 --:--:-- --:--:-- --:--:-- 47000
[upload] {"ok":true,"path":"/files/test.txt"}

[download] Downloading file...
[download]   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
[download]                                  Dload  Upload   Total   Spent    Left  Speed
100    13  100    13    0     0   5922      0 --:--:-- --:--:-- --:--:--  6500
[download] test content
```
