// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"

	types "github.com/gogo/protobuf/types"
)

func (dst *ApplicationLink) SetFields(src *ApplicationLink, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "default_formatters":
			if len(subs) > 0 {
				var newDst, newSrc *MessagePayloadFormatters
				if (src == nil || src.DefaultFormatters == nil) && dst.DefaultFormatters == nil {
					continue
				}
				if src != nil {
					newSrc = src.DefaultFormatters
				}
				if dst.DefaultFormatters != nil {
					newDst = dst.DefaultFormatters
				} else {
					newDst = &MessagePayloadFormatters{}
					dst.DefaultFormatters = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DefaultFormatters = src.DefaultFormatters
				} else {
					dst.DefaultFormatters = nil
				}
			}
		case "tls":
			if len(subs) > 0 {
				return fmt.Errorf("'tls' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TLS = src.TLS
			} else {
				var zero bool
				dst.TLS = zero
			}
		case "skip_payload_crypto":
			if len(subs) > 0 {
				return fmt.Errorf("'skip_payload_crypto' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SkipPayloadCrypto = src.SkipPayloadCrypto
			} else {
				dst.SkipPayloadCrypto = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationLinkRequest) SetFields(src *GetApplicationLinkRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				newDst = &dst.ApplicationIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetApplicationLinkRequest) SetFields(src *SetApplicationLinkRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				newDst = &dst.ApplicationIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "link":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationLink
				if src != nil {
					newSrc = &src.ApplicationLink
				}
				newDst = &dst.ApplicationLink
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationLink = src.ApplicationLink
				} else {
					var zero ApplicationLink
					dst.ApplicationLink = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationLinkStats) SetFields(src *ApplicationLinkStats, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "linked_at":
			if len(subs) > 0 {
				return fmt.Errorf("'linked_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LinkedAt = src.LinkedAt
			} else {
				dst.LinkedAt = nil
			}
		case "network_server_address":
			if len(subs) > 0 {
				return fmt.Errorf("'network_server_address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NetworkServerAddress = src.NetworkServerAddress
			} else {
				var zero string
				dst.NetworkServerAddress = zero
			}
		case "last_up_received_at":
			if len(subs) > 0 {
				return fmt.Errorf("'last_up_received_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LastUpReceivedAt = src.LastUpReceivedAt
			} else {
				dst.LastUpReceivedAt = nil
			}
		case "up_count":
			if len(subs) > 0 {
				return fmt.Errorf("'up_count' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpCount = src.UpCount
			} else {
				var zero uint64
				dst.UpCount = zero
			}
		case "last_downlink_forwarded_at":
			if len(subs) > 0 {
				return fmt.Errorf("'last_downlink_forwarded_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LastDownlinkForwardedAt = src.LastDownlinkForwardedAt
			} else {
				dst.LastDownlinkForwardedAt = nil
			}
		case "downlink_count":
			if len(subs) > 0 {
				return fmt.Errorf("'downlink_count' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DownlinkCount = src.DownlinkCount
			} else {
				var zero uint64
				dst.DownlinkCount = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AsConfiguration) SetFields(src *AsConfiguration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "pubsub":
			if len(subs) > 0 {
				var newDst, newSrc *AsConfiguration_PubSub
				if (src == nil || src.PubSub == nil) && dst.PubSub == nil {
					continue
				}
				if src != nil {
					newSrc = src.PubSub
				}
				if dst.PubSub != nil {
					newDst = dst.PubSub
				} else {
					newDst = &AsConfiguration_PubSub{}
					dst.PubSub = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.PubSub = src.PubSub
				} else {
					dst.PubSub = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetAsConfigurationRequest) SetFields(src *GetAsConfigurationRequest, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message GetAsConfigurationRequest has no fields, but paths %s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}

func (dst *GetAsConfigurationResponse) SetFields(src *GetAsConfigurationResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *AsConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &AsConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *NsAsHandleUplinkRequest) SetFields(src *NsAsHandleUplinkRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ups":
			if len(subs) > 0 {
				return fmt.Errorf("'application_ups' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationUps = src.ApplicationUps
			} else {
				dst.ApplicationUps = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EncodeDownlinkRequest) SetFields(src *EncodeDownlinkRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if (src == nil || src.VersionIds == nil) && dst.VersionIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.VersionIds
				}
				if dst.VersionIds != nil {
					newDst = dst.VersionIds
				} else {
					newDst = &EndDeviceVersionIdentifiers{}
					dst.VersionIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.VersionIds = src.VersionIds
				} else {
					dst.VersionIds = nil
				}
			}
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &ApplicationDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Parameter = src.Parameter
			} else {
				var zero string
				dst.Parameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EncodeDownlinkResponse) SetFields(src *EncodeDownlinkResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &ApplicationDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DecodeUplinkRequest) SetFields(src *DecodeUplinkRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if (src == nil || src.VersionIds == nil) && dst.VersionIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.VersionIds
				}
				if dst.VersionIds != nil {
					newDst = dst.VersionIds
				} else {
					newDst = &EndDeviceVersionIdentifiers{}
					dst.VersionIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.VersionIds = src.VersionIds
				} else {
					dst.VersionIds = nil
				}
			}
		case "uplink":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationUplink
				if (src == nil || src.Uplink == nil) && dst.Uplink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Uplink
				}
				if dst.Uplink != nil {
					newDst = dst.Uplink
				} else {
					newDst = &ApplicationUplink{}
					dst.Uplink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Uplink = src.Uplink
				} else {
					dst.Uplink = nil
				}
			}
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Parameter = src.Parameter
			} else {
				var zero string
				dst.Parameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DecodeUplinkResponse) SetFields(src *DecodeUplinkResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationUplink
				if (src == nil || src.Uplink == nil) && dst.Uplink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Uplink
				}
				if dst.Uplink != nil {
					newDst = dst.Uplink
				} else {
					newDst = &ApplicationUplink{}
					dst.Uplink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Uplink = src.Uplink
				} else {
					dst.Uplink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DecodeDownlinkRequest) SetFields(src *DecodeDownlinkRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "version_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceVersionIdentifiers
				if (src == nil || src.VersionIds == nil) && dst.VersionIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.VersionIds
				}
				if dst.VersionIds != nil {
					newDst = dst.VersionIds
				} else {
					newDst = &EndDeviceVersionIdentifiers{}
					dst.VersionIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.VersionIds = src.VersionIds
				} else {
					dst.VersionIds = nil
				}
			}
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &ApplicationDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}
		case "formatter":
			if len(subs) > 0 {
				return fmt.Errorf("'formatter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formatter = src.Formatter
			} else {
				var zero PayloadFormatter
				dst.Formatter = zero
			}
		case "parameter":
			if len(subs) > 0 {
				return fmt.Errorf("'parameter' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Parameter = src.Parameter
			} else {
				var zero string
				dst.Parameter = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DecodeDownlinkResponse) SetFields(src *DecodeDownlinkResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &ApplicationDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AsConfiguration_PubSub) SetFields(src *AsConfiguration_PubSub, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "providers":
			if len(subs) > 0 {
				var newDst, newSrc *AsConfiguration_PubSub_Providers
				if (src == nil || src.Providers == nil) && dst.Providers == nil {
					continue
				}
				if src != nil {
					newSrc = src.Providers
				}
				if dst.Providers != nil {
					newDst = dst.Providers
				} else {
					newDst = &AsConfiguration_PubSub_Providers{}
					dst.Providers = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Providers = src.Providers
				} else {
					dst.Providers = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AsConfiguration_PubSub_Providers) SetFields(src *AsConfiguration_PubSub_Providers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "mqtt":
			if len(subs) > 0 {
				return fmt.Errorf("'mqtt' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MQTT = src.MQTT
			} else {
				var zero AsConfiguration_PubSub_Providers_Status
				dst.MQTT = zero
			}
		case "nats":
			if len(subs) > 0 {
				return fmt.Errorf("'nats' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NATS = src.NATS
			} else {
				var zero AsConfiguration_PubSub_Providers_Status
				dst.NATS = zero
			}
		case "aws_iot":
			if len(subs) > 0 {
				return fmt.Errorf("'aws_iot' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AWSIoT = src.AWSIoT
			} else {
				var zero AsConfiguration_PubSub_Providers_Status
				dst.AWSIoT = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}