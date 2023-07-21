
config:
clientServer:
  name: stats-client
  grpcServerName: stats-server
  grpcGroupName: DEFAULT_GROUP
  grpcClusterName: DEFAULT
  port: 8032 
jaeger:
  serviceName: stats-client
  hostPort: localhost:6831
db:
  name: "rec_visual"
  address: "localhost:3306"
  userName: root
  password: 123456
  maxIdleConnects: 10
  maxOpenConnects: 100
  connMaxLifetimeHour: 1
redis:
  address: "localhost:6379"
  userName: root
  password: 123456
  database: 0