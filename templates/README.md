导出`yaml`文件：

```sh
$ kubectl get deployments --all-namespaces=true --export -o yaml > demo.yaml #导出
$ kubectl run busybox --image=busybox  --dry-run -o yaml > demo.yaml #dry run
```

