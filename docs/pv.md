# pv
```yaml
kind: PersistentVolume
apiVersion: v1
metadata:
  name: name_of_pv
  namespace: demo
  labels:
    name: name_of_pv
spec:
  capacity:
    storage: 100Mi
  nfs:
    server: 127.0.0.1
    path: /data/xx-pv
  accessModes:
  - ReadWriteOnce
```

# pvc 
```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: name_of_pvc
  namespace: demo
spec:
  accessModes:
  - ReadWriteOnce
  selector:
    matchLabels:
      name: name_of_pv
  resources:
    requests:
      storage: 100Mi
```