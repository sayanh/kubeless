# Instructions

### Kubecontroller isolation
- Start minikube with RBAC and enabled storage class

  ```minikube start --memory=4096 --kubernetes-version v1.8.5 --bootstrapper kubeadm```

- Create two namespaces `kubeless` and `kubeless-new`

  ```kubectl create namespace kubeless```

- Deploy kubeless as follows:

    ```kubectl apply -f kubeless-rbac-v0.3.0.yaml -n kubeless```

    ```kubectl apply -f kubeless-rbac-v0.3.0-diff-namespace.yaml -n kubeless-new```

- Check `kubeless-contoller` logs in each namespace to ensure that it is listening to the namespace where it is deployed

  ```kubectl log po/kubeless-controller-9748799d5-c7nfw -n kubeless-new```

  > msg=+++++++++++++++++NAMESPACE++++++++++++++++++++++
  time="2017-12-11T14:39:29Z" level=info msg=+++++++++++++++++++kubeless-new++++++++++++++++++++
  time="2017-12-11T14:39:29Z" level=info msg="Starting kubeless controller" pkg=controller
  time="2017-12-11T14:39:29Z" level=info msg="Kubeless controller synced and ready" pkg=controller```

### Kubeless client isolation
- Create two users and configure the keys as described [here][1]
- Here we created users `employee` and `employee-new`
- Create confined RBAC policies as follows for the above users:

  ```kubectl apply -f crd-role-employee.yaml```

  ```kubectl apply -f crd-role-employee-new.yaml```

- To verify access control

  ```kubectl auth can-i list function --namespace kubeless --as employee```
  > yes

- Use context `employee-context` which binds to user `employee` or `employee-context-new` bind to user `employee-new`

  ```kubectl config use-context employee-context```

- Download kubeless cli from [here][2]
- Since kubeless uses local kubeconfig, CRUD operations for functions can be done as below:
  E.g.
  ```kubeless function list --namespace kubeless-new```
  > FATA[0000] functions.k8s.io is forbidden: User "employee" cannot list functions.k8s.io in the namespace "kubeless-new"

  ```kubeless function list --namespace kubeless```
  >NAME  	NAMESPACE	HANDLER  	RUNTIME	TYPE	TOPIC	DEPENDENCIES	STATUS

  >testjs	kubeless 	hello.foo	nodejs8	HTTP	     	            	1/1 READY

[1]: https://docs.bitnami.com/kubernetes/how-to/configure-rbac-in-your-kubernetes-cluster/#use-case-1-create-user-with-limited-namespace-access

[2]: https://github.com/kubeless/homebrew-tap
