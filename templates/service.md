# service模板
```yaml
apiVersion: v1
kind: Service
metadata:
  name: ping-pong-service # Service名称
spec:
  selector:
    app: ping-pong # 选择labels
  type: ClusterIP # NodePort、ClusterIP、LoadBalancer
  ports:
  - protocol: TCP # TCP、UDP、
    port: 8080 # Service的端口
    targetPort: 80 # 指定Pod中容器的端口，可以是一个string类型的名字，对应的真实端口值由各个后端Pod定义，这样同一组Pod无须使用同一个Port，更加灵活
    #nodePort: 5050 # NodePort类型可以指定映射的宿主机端口，范围在30000-32767，如果没有指定则随机分配
```