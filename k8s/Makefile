.PHONY: all namespace configmap secrets db app ingress vault deploy status clean

VAULT_NAMESPACE=vault-ns
VAULT_VERSION=0.29.1
VAULT_TOKEN=YourToken  # Replace with your actual Vault token


all: deploy status

vault-token:
	kubectl create secret generic vault-token -n external-secret-ns --from-literal=token="$(VAULT_TOKEN)" # --dry-run=client -o yaml | kubectl apply -f -

namespace:
	kubectl apply -f namespace/namespace.yaml

configmap:
	kubectl apply -f configmap.yaml

secrets:
	kubectl apply -f eso/cluster-secretstore.yaml
	kubectl apply -f eso/external-secret.yaml

db:
	kubectl apply -f postgres/postgres-deployment.yaml
	kubectl apply -f postgres/postgres-service.yaml

app:
	kubectl apply -f go-crud-app/go-crud-app-deployment.yaml
	kubectl apply -f go-crud-app/go-crud-app-service.yaml

ingress:
	kubectl apply -f ingress.yaml

vault:
	helm repo add hashicorp https://helm.releases.hashicorp.com
	helm repo update
	helm install vault hashicorp/vault --namespace $(VAULT_NAMESPACE) --version $(VAULT_VERSION)

deploy: namespace vault configmap secrets db app ingress
	@echo "✅ Deployment complete!"

status:
	kubectl get all -A

clean:
	helm uninstall vault --namespace $(VAULT_NAMESPACE)
	kubectl delete namespace $(VAULT_NAMESPACE)
	kubectl delete -f ingress.yaml
	kubectl delete -f go-crud-app/go-crud-app-deployment.yaml
	kubectl delete -f go-crud-app/go-crud-app-service.yaml
	kubectl delete -f postgres/postgres-deployment.yaml
	kubectl delete -f postgres/postgres-service.yaml
	kubectl delete -f eso/external-secret.yaml
	kubectl delete -f eso/cluster-secretstore.yaml
	kubectl delete -f configmap.yaml
	@echo "🗑️  Cleanup complete!"
