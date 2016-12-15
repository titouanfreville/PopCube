# Projet de fin d'année SUPINFO M Sc 1 2016-2017

*Ce document récapitule le workflow a suivre pour le projet de fin d'année. Merci de respecter ces étapes.*

**Les branches DEMO et DEVELOP doivent rester stable en toute circonstance. Si un merge casse la branche, il doit être revert puis corriger avant de le merge de nouveau.**

**La branche master sera mise à jour toute les semaines sur la branche develop. Aucun autre merge ne doit être réaliser sur master en aucune circonstance**

**La branche demo doit êter mise à jour sur develop avant chaque préparation de demo ou en fin de semaine**

## Préparer son environement de travail 

- Installez `git`, `docker` et `docker-compose`
- Récupérer le code source du projet : `git clone git@github.com:titouanfreville/SUPINFO-MSC1-PJT.git || git clone https://github.com/titouanfreville/SUPINFO-MSC1-PJT.git`
- Aller dans le dossier obtenu
- `docker-compose up`

## Proposer une feature / Séparer les tâches

- Créer une issue sur github contenant un nom (résumant le travail à faire) et l'intégralité de la feature (ex : Nom : HELLO WORLD, Contenu : Créer une page de l'application permettant d'afficher le message : 'Hello World')
- Signaler la création de l'application.
- Voter et ce mettre d'accord ;)

## Travailler sur une issue

**Quand une feature est acceptée**

- Créer une branche depuis la branche 'dévelop' nommer `issue-Numéro_de_l'issue-Nom`
- Créer les test unitaires liés à l'issue si non existant
- Faite votre feature
- Tester (lancer les test unitaires, etc.)
- Ouvrir une pull request de votre branche vers master

### Format de commit 

- ajouter le template de commit a la configuration git. `git config commit.template {Clone_ROOT}/.git_commit_message.txt`
- Un commit doit contenir : le numéro de l'issue (requis), un résumé du travail réaliser(requis), un détails du travail réalisé, la liste de ce qu'il reste à faire, un ping sur les différentes personne concerné par l'issue.

## Finir une issue 

**Quand une feature est en Pull Request**

- Une personne différente de celle ayant fait l'issue vérifie le code
- La feature est testé et une démo doit être faite (merge sur la branche demo)
- La PR est acceptée
- Merge la branche dans develop
- Tester que develop reste stable