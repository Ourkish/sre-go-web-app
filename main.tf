provider "kubernetes" {
  # Ensure that kubectl is using minikube context.
  # You can verify it using: `kubectl config current-context`
  config_path = "/home/debian/.kube/config"
}

provider "helm" {
  kubernetes {
    config_path = "/home/debian/.kube/config"
  }
}

# Deploying Web app 
resource "kubernetes_deployment" "sre_hiring_app" {
  metadata {
    name = "sre-hiring-app-deployment"
    labels = {
      app = "sre-hiring"
    }
  }

  spec {
    replicas = 3

    selector {
      match_labels = {
        app = "sre-hiring"
      }
    }

    template {
      metadata {
        labels = {
          app = "sre-hiring"
        }
        annotations = {
          "prometheus.io/scrape" = "true"
          "prometheus.io/port"   = "8080"
          "prometheus.io/path"   = "/metrics"
        }
      }

      spec {
        container {
          image = "sre-hiring:local"
          name  = "sre-hiring"

          port {
            container_port = 8080
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "sre_hiring_service" {
  metadata {
    name = "sre-hiring-service"
  }

  spec {
    selector = {
      app = "sre-hiring"
    }

    port {
      port        = 8080
      target_port = 8080
    }

    type = "NodePort"
  }
}

resource "helm_release" "prometheus" {
  name       = "prometheus"
  repository = "https://prometheus-community.github.io/helm-charts"
  chart      = "prometheus"
}

resource "helm_release" "grafana" {
  name       = "grafana"
  repository = "https://grafana.github.io/helm-charts"
  chart      = "grafana"

  set {
    name  = "adminPassword"
    value = "Grafana@OVH#"
  }

  set {
    name  = "persistence.enabled"
    value = "true"
  }

  set {
    name  = "persistence.size"
    value = "10Gi"
  }
}

