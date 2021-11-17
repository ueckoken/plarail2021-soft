// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/proto/statesync.proto#L28-L85
export type StationId =
  | "unknown"
  | "motoyawata_s1"
  | "motoyawata_s2"
  | "motoyawata_s3"
  | "motoyawata_s4"
  | "motoyawata_s5"
  | "motoyawata_s6"
  | "iwamotocho_s1"
  | "iwamotocho_s2"
  | "iwamotocho_s4"
  | "iwamotocho_b1"
  | "iwamotocho_b2"
  | "iwamotocho_b3"
  | "iwamotocho_b4"
  | "kudanshita_s5"
  | "kudanshita_s6"
  | "sasazuka_b1"
  | "sasazuka_b2"
  | "sasazuka_s1"
  | "sasazuka_s2"
  | "sasazuka_s3"
  | "sasazuka_s4"
  | "sasazuka_s5"
  | "meidaimae_s1"
  | "meidaimae_s2"
  | "chofu_s1"
  | "chofu_s2"
  | "chofu_s3"
  | "chofu_s4"
  | "chofu_s5"
  | "chofu_s6"
  | "chofu_b1"
  | "chofu_b2"
  | "chofu_b3"
  | "chofu_b4"
  | "chofu_b5"
  | "kitano_b1"
  | "kitano_b2"
  | "kitano_b3"
  | "kitano_b4"
  | "kitano_s1"
  | "kitano_s2"
  | "kitano_s3"
  | "kitano_s4"
  | "kitano_s5"
  | "kitano_s6"
  | "takao_s1"
  | "takao_s2"

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/proto/statesync.proto#L10-L14
export type StationState = "UNKNOWN" | "ON" | "OFF"

export type StationMessage = {
  station_name: StationId
  state: StationState
}

// https://github.com/ueckoken/plarail2021-soft/blob/a3982c4ef4b20e371052b4ad36b777a04ed67d1a/backend/external/pkg/clientHandler/clientHandler.go#L28-L31
export type Message = StationMessage
