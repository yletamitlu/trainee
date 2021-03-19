# trainee


## API методы:
- Метод сохранения статистики
- Метод показа статистики
- Метод сброса статистики

```
http://localhost:80/save – запрос на сохранение статистики
http://localhost:80/get – запрос на показ всей статистики
http://localhost:80/clear – запрос на удаление всей статистики
```

### Запрос на сохранение статистики
Метод: POST
```
Путь: http://localhost:80/save
Пример body:
{
    "date": "1992-08-15",
    "views": 2,
    "clicks": 11,
    "cost": 354.56
}
```

### Запрос на показ всей статистики
Метод: GET
```
Путь: http://localhost:80/get?since=from&until=to

Пример: http://localhost:80/get?since=1992-08-15&until=2010-08-15
```
Можно указать параметр, по которому можно сортировать:
```
Пример: http://localhost:80/get?since=1960-08-15&until=2010-08-15&param=cost
 ```
### Запрос сброса статистики
Метод: POST
```
Путь: http://localhost:80/clear
```
## Запуск 
```
sudo docker build -t stat . 
sudo docker run -p 80:80 --name stat -t stat
```
