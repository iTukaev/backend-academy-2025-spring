package tbank.ab.config

import cats.effect.IO
import pureconfig.{ConfigReader, ConfigSource}

final case class AppConfig()(using
  val server: ServerConfig,
  val auth: AuthConfig,
  val database: DbConfig,
  val s3: S3Config,
  val redis: RedisConfig
)

object AppConfig {
  private case class AppConfigView(
    server: ServerConfig,
    auth: AuthConfig,
    database: DbConfig,
    s3: S3Config,
    redis: RedisConfig
  ) derives ConfigReader

  def load(source: ConfigSource): IO[AppConfig] =
    IO.delay(source.loadOrThrow[AppConfigView])
      .map(view =>
        AppConfig()(using
          server = view.server,
          auth = view.auth,
          database = view.database,
          s3 = view.s3,
          redis = view.redis
        )
      )
}
