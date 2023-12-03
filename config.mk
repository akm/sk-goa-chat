# STAGE
# - local
# - github
# - staging
# - production

# PROCESS_TYPE        | STAGE                    | apisvr      | ui     | mysql    | firebase authentication | メモ
# --------------------|--------------------------|-------------|--------|----------|-------------------------|--------------
# server              | staging,production       | 8000,8080   | 4173   | (3306)   | (9099)                  |
# e2e test            | staging,production       | (8000,8080) | (4173) | (3306)   | (9099)                  | 対象のサーバーにアクセスするので特にサーバーを起動したりはしない
# server              | local                    | 8000,8080   | 5173   | 3306     | 9099
# e2e test            | local,github             | 8001,8081   | 4173   | 3307     | 9090
# unit test / apisvr  | local,github             | -,-         | -      | 3311 + N | 9091 + N
# unit test / ui      | local,github             | -,-         | -      | -        | -

# local以外では環境変数 STAGE を設定し、変数 STAGE のデフォルト値を上書きします。
# STAGE=github
# STAGE=staging
# STAGE=production
STAGE?=local

MYSQL_DATABASE_NAME=sk_goa_chat_db
