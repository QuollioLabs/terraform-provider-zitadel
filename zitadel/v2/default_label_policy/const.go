package default_label_policy

const (
	primaryColorVar        = "primary_color"
	hideLoginNameSuffixVar = "hide_login_name_suffix"
	warnColorVar           = "warn_color"
	backgroundColorVar     = "background_color"
	fontColorVar           = "font_color"
	primaryColorDarkVar    = "primary_color_dark"
	backgroundColorDarkVar = "background_color_dark"
	warnColorDarkVar       = "warn_color_dark"
	fontColorDarkVar       = "font_color_dark"
	disableWatermarkVar    = "disable_watermark"
	logoPathVar            = "logo_path"
	logoHashVar            = "logo_hash"
	logoURLVar             = "logo_url"
	iconPathVar            = "icon_path"
	iconHashVar            = "icon_hash"
	iconURLVar             = "icon_url"
	logoDarkPathVar        = "logo_dark_path"
	logoDarkHashVar        = "logo_dark_hash"
	logoURLDarkVar         = "logo_url_dark"
	iconDarkPathVar        = "icon_dark_path"
	iconDarkHashVar        = "icon_dark_hash"
	iconURLDarkVar         = "icon_url_dark"
	fontPathVar            = "font_path"
	fontHashVar            = "font_hash"
	fontURLVar             = "font_url"
	setActiveVar           = "set_active"
)

const (
	assetAPI       = "/assets/v1"
	labelPolicyURL = "/instance/policy/label"
	logoURL        = assetAPI + labelPolicyURL + "/logo"
	logoDarkURL    = logoURL + "/dark"
	iconURL        = assetAPI + labelPolicyURL + "/icon"
	iconDarkURL    = iconURL + "/dark"
	fontURL        = assetAPI + labelPolicyURL + "/font"
)
