// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkplatformtest

type TestGroup struct {
	Tag2 *Tag2
	Tag3 *Tag3

	sdkConfiguration sdkConfiguration
}

func newTestGroup(sdkConfig sdkConfiguration) *TestGroup {
	return &TestGroup{
		sdkConfiguration: sdkConfig,
		Tag2:             newTag2(sdkConfig),
		Tag3:             newTag3(sdkConfig),
	}
}
