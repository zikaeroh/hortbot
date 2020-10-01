module github.com/hortbot/hortbot

go 1.15

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/integrations/ocsql v0.1.6
	github.com/99designs/gqlgen v0.13.0
	github.com/alicebob/miniredis/v2 v2.13.3
	github.com/antchfx/htmlquery v1.2.3
	github.com/araddon/dateparse v0.0.0-20200930201622-fcfe3a02eb30
	github.com/bmatcuk/doublestar/v2 v2.0.1
	github.com/dghubble/trie v0.0.0-20200716043226-5a94efb202d5
	github.com/dustin/go-humanize v1.0.0
	github.com/ericlagergren/decimal v0.0.0-20191206042408-88212e6cfca9 // indirect
	github.com/felixge/httpsnoop v1.0.1
	github.com/fortytw2/leaktest v1.3.0
	github.com/friendsofgo/errors v0.9.2
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-redis/redis/v8 v8.2.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobuffalo/flect v0.2.2
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.13.0
	github.com/google/go-cmp v0.5.2
	github.com/gorilla/sessions v1.2.1
	github.com/goware/urlx v0.3.1
	github.com/hako/durafmt v0.0.0-20200710122514-c0fb7b4da026
	github.com/jackc/pgconn v1.7.0
	github.com/jackc/pgx/v4 v4.9.0
	github.com/jakebailey/irc v0.0.0-20190904051515-2d11e69506b0
	github.com/jarcoal/httpmock v1.0.6
	github.com/jessevdk/go-flags v1.4.1-0.20181221193153-c0795c8afcf4
	github.com/jmoiron/sqlx v1.2.1-0.20190826204134-d7d95172beb5
	github.com/joho/godotenv v1.3.0
	github.com/leononame/clock v0.1.6
	github.com/markbates/pkger v0.17.1
	github.com/maxbrunsfeld/counterfeiter/v6 v6.2.3
	github.com/nsqio/go-nsq v1.0.8
	github.com/ory/dockertest/v3 v3.6.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/posener/ctxutil v1.0.0
	github.com/prometheus/client_golang v1.7.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.2.1
	github.com/tomwright/queryparam/v4 v4.1.0
	github.com/valyala/quicktemplate v1.6.3
	github.com/vektah/gqlparser/v2 v2.1.0
	github.com/volatiletech/null/v8 v8.1.0
	github.com/volatiletech/sqlboiler/v4 v4.2.0
	github.com/volatiletech/strmangle v0.0.1
	github.com/wader/filtertransport v0.0.0-20200316221534-bdd9e61eee78
	github.com/zikaeroh/ctxjoin v0.0.0-20200613235025-e3d47af29310
	github.com/zikaeroh/ctxlog v0.0.0-20200613043947-8791c8613223
	go.deanishe.net/fuzzy v1.0.0
	go.opencensus.io v0.22.4
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20200930145003-4acb6c075d10
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sync v0.0.0-20200930132711-30421366ff76
	golang.org/x/tools v0.0.0-20200930213115-e57f6d466a48
	gotest.tools/v3 v3.0.2
	mvdan.cc/xurls/v2 v2.2.0
)

replace github.com/araddon/dateparse => github.com/zikaeroh/dateparse v0.0.0-20201001042724-56dec34db17c
