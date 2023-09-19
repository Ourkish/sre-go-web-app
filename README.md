# sre-web-app

## Objective

- Run a provided sre-web-app application
- Implement a deployment system with Kubernetes (minikub, k3d ... ) and Helm.
- You can improve, propose and implement everything you want.
- Monitor the application.

## Prerequisites

- Linux workstation (I am using Ubuntu 20.04) with at least 15 Gb ram (r2-15 flavor on OVH) .
  `````
  cat /etc/*release*
  `````
- Install the following packages [curl, Docker, Minikube, kubectl, Helm, Terraform]:
  `````
  sudo mkdir /opt/sre-web-app
  sudo cd /opt/sre-web-app
  git clone https://github.com/Ourkish/sre-web-app.git
  sudo chmod +x auto_install.sh
  sudo ./auto_install.sh
  `````
- get the web app
  in this lab we used a Java Web app (https://hub.docker.com/r/tkgregory/sample-metrics-application)

- Ensure your Horizon account is prepared, then proceed to the 'Network' section. From there, access 'Security Group' and establish fresh rules that permit TCP, HTTP, and Ingress on 
  ports 3000, 9090, and 30080. Set the authorization to '0.0.0.0/0' to enable browser-based checks on the instance's public IP.
  
## Understanding the operation

- Minikube provides a local Kubernetes environment where you can deploy and manage containerized applications using Kubernetes resources.
- Docker containers can be built and pushed to container registries. Kubernetes can then pull these containers and deploy them as pods (the smallest deployable units in Kubernetes).
- kubectl is used to interact with the Minikube Kubernetes cluster. You can create pods, services, deployments, and other Kubernetes resources using kubectl commands.
- Helm simplifies the installation of complex applications in Kubernetes. Helm charts describe how to deploy applications, including the necessary Kubernetes resources. You can use 
  Helm to install and manage applications 
  in your Minikube cluster.
- Terraform can manage Kubernetes resources in Minikube as well as other infrastructure resources outside of Kubernetes. It uses Terraform providers to interact with the Kubernetes API 
  and provision resources like 
  namespaces, secrets, and ConfigMaps.

## In Practice

- Quick Start :
  Clone the project on your environement and run those commands (if not already done !) :
  `````
  git clone https://github.com/Ourkish/sre-web-app.git
  sudo chmod +x dep_install.sh
  sudo ./auto_install.sh
  sudo terraform apply
  `````
  
- Allow EST/WEST traffic on exposed ports :
  `````
  kubectl port-forward --address 141.94.106.42 service/grafana 3000:80 > /dev/null 2>&1
  kubectl port-forward --address 141.94.106.42 service/prometheus-server 9090:80 > /dev/null 2>&1 &
  kubectl port-forward --address 141.94.106.42 service/sample-metrics-app-service 30080:8080 > /dev/null 2>&1 &
  `````

- Add a cron to check disk and to be notified by email if disk run's out of space :
  `````
  sudo vim /etc/crontab
  0 * * * * /usr/bin/go run /home/ubuntu/sre-web-app/HealthChecks.go >> /var/log/HealthChecks.log 2>&1
  `````

  you can configure your creds for SMTP service here :
  `````
  sudo vim /home/ubuntu/sre-web-app/HealthChecks.go
  recipient := "your.email@example.com"
  smtpServer := "smtp.example.com"
  smtpPort := "587"
  smtpUsername := "your.smtp.username"
  smtpPassword := "your.smtp.password"
  `````
  
- Below are screenshots showcasing the results following the deployment and execution of the code.

  Deploy using Helm and Terraform :

  ![deploy](https://github.com/Ourkish/sre-web-app/assets/67292535/9b5cdf6d-1465-459c-8f42-d3e8b31b8ec5)

  
  Application Monitoring :
  
  ![jvm_dashboard](https://github.com/Ourkish/sre-web-app/assets/67292535/c888760b-1166-471e-925f-882cb5fda8b5)

  Deployment using terraform :
  
  ![pods](https://github.com/Ourkish/sre-web-app/assets/67292535/fe3d8bcc-05ce-4c50-ad74-7d34f094002d)

  App metrics after deployment :
  
  ![web_app_metrics](https://github.com/Ourkish/sre-web-app/assets/67292535/b89fd98e-b968-4d83-bb5d-7dee749bc0f7)

  Kubernetes Dashboard :
  
  ![kube_dashboard](https://github.com/Ourkish/sre-web-app/assets/67292535/8762a72c-2ac2-47e4-a965-baf36bc6692a)

  To destroy the cluster use this command :
  `````
  sudo terraform destroy
  `````
  you will be asked to confirme, if you agree to destroy you should see an output like this :
  
  ![destroy_cluster](https://github.com/Ourkish/sre-web-app/assets/67292535/f59a86f6-ebb8-4583-b93a-82ba823591c6)



  
  
