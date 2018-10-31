# job

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  namespace: yangmc-test
  name: init-yapi-db
spec:
  template:
    spec:
      containers:
      - name: init-yapi-db
        image: registry.cn-qingdao.aliyuncs.com/yangmc/kst-yapi:1.0.0
        command: ["node", "server/install.js"]
      restartPolicy: Never #Never
  backoffLimit: 4

```

