### ArgoCD example
  
Simple ArgoCD usage example.  
ArgoCD can be set up on minikube cluster by installing it with helm (through terraform in respective folder) or it can be installed on any different cluster (EKS for example) but helm provider needs to be changed properly. 
   
Added example frontend and backend app with their Dockerfiles along with kuberentes manifests files with their deployments and ingress.  
  
Github workflow is set up to build and push backend and frontend images with a new tag to the docker hub on push to the main branch (regardless if any change was made in the actual apps - for easier showcase). Then it automatically changes kubernetes deployment manifests to use images with newly created tags (change and auto-commit during workflow). When workflow is completed, created drift triggers ArgoCD to deploy new version of the app.   