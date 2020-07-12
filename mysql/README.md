### Kubectl commands

kubectl create secret generic mysql-pass --from-literal=password='12345678'

kubectl apply -f service.yaml   

kubectl apply -f persistent-volume.yaml

kubectl apply -f deployment.yaml


kubectl get pods

kubectl get services

kubectl get persistentvolumeclaim

kubectl get deployments 

kubectl delete deployments --all


### Running

kubectl exec -it POD_NAME bash

mysql -uroot -p