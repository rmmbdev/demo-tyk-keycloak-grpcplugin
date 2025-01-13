# This is a sample implementation of integrating Tyk and Keycloak with JWT authentication

Run these commands sequentially 

build fatapi image
``` bash
cd fastapi && docker build --network=host -t fastapi-demo . 
```

back to root 
``` bash
cd ..
```

run containers
``` bash
docker compose up
```

install `busybox` 

``` bash
sudo pacman -S busybox
```

copy `busybox` to realted docker (this is because of distroless images)
``` bash
docker cp /usr/bin/busybox demo-tyk-keyclock-keycloak-1:/busybox 
``` 

### To import realm configs do this
One of the most important parts of integrating these two is configuring keycloak for providing suitable jwt. But we can use the exported keycloak realm which is configured earlier.

copy config to docker
``` bash
docker cp ./my-realm.json demo-tyk-keycloak-keycloak-q:/opt/keycloak/data
```

restore 
``` bash
docker exec demo-tyk-keyclock-keycloak-1 /opt/keycloak/bin/kc.sh import --file=/opt/keycloak/data/my-realm.json
```

### create api
``` bash
curl --request POST \
  --url http://localhost:8090/tyk/apis/oas \
  --header 'Content-Type: application/json' \
  --header 'X-Tyk-Authorization: foo' \
  --data 'add contents of sample.api-oas.json here'
```

### create policy
``` bash 
curl --request POST \
  --url http://localhost:8090/tyk/policies \
  --header 'Content-Type: application/json' \
  --header 'x-tyk-authorization: foo' \
  --data 'add contents of sample.policy.json here'

```

### test the pipeline

get token
``` bash
curl --request POST \
  --url http://localhost:8100/realms/my-realm/protocol/openid-connect/token \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data client_id=conf-client \
  --data grant_type=password \
  --data scope=openid \
  --data username=ahmad \
  --data password=123 \
  --data client_secret=ASDN46a1SJ4PetQqfCbnVxn8oO3STugV \
  --data redirect_url=http://tyk:8080/auth/oidc
```

test api
``` bash
curl --request GET \
  --url http://localhost:8090/api/items/1 \
  --header 'Authorization: Bearer <accecc-token>'
```


### Debug
attach bash to container for test purposes
``` bash
docker exec -it demo-tyk-keyclock-keycloak-1 /busybox sh
```