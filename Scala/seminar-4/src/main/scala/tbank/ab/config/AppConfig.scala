package tbank.ab.config

import cats.effect.IO
import pureconfig.ConfigConvert.catchReadError
import pureconfig.configurable.genericMapReader
import pureconfig.{ConfigObjectSource, ConfigReader, ConfigSource}
import tbank.ab.domain.animal.{AnimalId, AnimalInfo}

final case class AppConfig(
  port: Int,
  auth: AuthConfig,
  animals: Map[AnimalId, AnimalInfo]
) derives ConfigReader

object AppConfig:
  def load(source: ConfigSource): IO[AppConfig] =
    IO.delay(source.loadOrThrow[AppConfig])

  given ConfigReader[Map[AnimalId, AnimalInfo]] =
    genericMapReader[AnimalId, AnimalInfo](
      catchReadError[AnimalId](id => AnimalId(id))
    )
