# SBD Exercise 7
In today's lab we are going to setup our own small cluster.
You will be required to work in groups of 2+ people.

- [ ] Connect your computers to the same network 
  - The network should not restrict any communication, it's best to use your phone's hotspot
- [ ] Init Docker Swarm Mode
  - Setup one node to act as a manager
  - Other nodes should join as workers
- [ ] Adapt Traefik configuration to work in Docker Swarm
  - No SSL / HTTPS needed!
- [ ] Deploy the stack to your swarm
  - Replicate SWS and the Orderservice to every node (global)
  - Deploy one replica of Traefik to the manager
  - Deploy one replica of the Database to a single node in the swarm
  - Deploy one replica of Minio to a single node in the swarm
- [ ] Any host on the network should be able to reach Traefik hosted on the manager node
  - The frontend should be accessible at http://<manager-node-ip>
  - The API should be accessible at http://orders.<manager-node-ip>
- [ ] Remove usernames and passwords from environment variables and put them into Docker Secrets
- [ ] Keep track of all commands you used and add them to the `setup.md` file
In the end your Swarm should look something like this:
```
                  3 Node Swarm Setup                  
                                                                                    
┌──────────────┐   ┌──────────────┐   ┌──────────────┐
│    Node 1    │   │    Node 2    │   │    Node 3    │
├──────────────┤   ├──────────────┤   ├──────────────┤
│  Manager 1   │   │   Worker 1   │   │   Worker 2   │
├──────────────┤   ├──────────────┤   ├──────────────┤
│   Traefik    │   │  PostgresDb  │   │    Minio     │
│     SWS      │   │     SWS      │   │     SWS      │
│ Orderservice │   │ Orderservice │   │ Orderservice │
└──────────────┘   └──────────────┘   └──────────────┘ 
```

## Tips and Tricks
Have a look at the following resources to understand how to set up Traefik on a 
Docker Swarm cluster. **DO NOT** set up self-signed certificates or Letsencrypt!
- https://doc.traefik.io/traefik/setup/swarm/#create-a-dockercomposeswarmyaml

Binding a container to a specific host can be archived with the following deployment
specification:
```yml
deploy:
  replicas: 1
  placement:
    constraints:
      # bind to a specific hostname
      - node.hostname==<target-pc-hostname>
      # OR bind to manager node
      - node.role == manager
```

How to deploy and setup a Docker Swarm cluster:
- https://docs.docker.com/engine/swarm/swarm-tutorial/

Adding secrets to the docker-compose file:
- https://docs.docker.com/compose/how-tos/use-secrets/

Beware that service labels in Docker Swarm need to be placed after the deploy section!
```yml
deploy:
  labels:
    - traefik.enable=true
    ...
```

## Docker Swarm Commands
- `docker swarm init` Inits a swarm on the current node
- `docker swarm join ...` Joins an already existing swarm
- `docker node ps` Lists all nodes connected to the swarm
- `docker service ls` Lists all services running on the swarm
- `docker service inspect --pretty <service>` Inspect a service
- `docker stack deploy -c <docker-compose.swarm.yml> <name-of-stack>` Deploy a stack to the cluster OR update current cluster with new configuration
- `docker stack rm <name-of-stack>` Delete stack, equivalent to docker-compose down
- `docker service ps <service-name> (--no-trunc)`  List tasks services, useful to debug failing services