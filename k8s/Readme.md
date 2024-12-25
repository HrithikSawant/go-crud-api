# Go CRUD Application - Kubernetes Deployment

This project contains Kubernetes manifests for deploying a Go-based CRUD application and its PostgreSQL database. The deployment is designed to work in a Kubernetes cluster with logical node separation and includes support for ingress and secure environment variable management.

## Project Overview

### Application Components
1. **Go CRUD App**: The primary application to handle CRUD operations.
2. **PostgreSQL Database**: Backend database for data storage.

### Kubernetes Features Used
- ConfigMaps for non-sensitive configurations
- Secrets for sensitive data
- Deployments for managing pods
- Services for exposing the application and database
- Ingress for HTTP routing

## Prerequisites

Before proceeding, ensure you have the following:

1. A running **Kubernetes cluster** (e.g., Minikube, Docker Desktop, or a cloud provider).
2. **kubectl** installed and configured to interact with the cluster.
3. An **Ingress Controller** deployed in your cluster (e.g., Nginx Ingress).
4. Basic knowledge of Kubernetes concepts (Deployments, Services, ConfigMaps, Secrets, and Ingress).

## Setup Instructions

### 1. Start Kubernetes Cluster
For local testing with Minikube, use the following command:

```bash
minikube start --driver=virtualbox
```
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```

Replace `<node-name>` with your actual node names.

### 2. Configure Secrets
Generate base64-encoded values for database credentials:

```bash
echo -n "Your-User" | base64
echo -n "Your-Pass" | base64
```

Update `secret.yaml` with the encoded values for `DB_USER` and `DB_PASSWORD`.

### 3. Deploy Resources
Use the provided script to deploy all components:

```bash
bash deploy.sh
```

The script applies the ConfigMap, Secret, Deployments, Services, and Ingress in sequence.

### 4. Verify Deployment
Check the status of the resources:

```bash
kubectl get pods
kubectl get services
kubectl get ingress
```

### 5. Access the Application
1. Use Minikube tunneling to expose ingress:

   ```bash
   minikube tunnel
   ```

2. Access the application at the host specified in `ingress.yaml` (e.g., `go-crud-app.local`). Ensure your `/etc/hosts` file is updated to map this host to the Minikube IP.

## Cleanup

To delete all deployed resources:

```bash
kubectl delete -f .
```

## License

This project is licensed under the MIT License.

## Support

For questions or issues, please raise a GitHub issue or contact the maintainer.

