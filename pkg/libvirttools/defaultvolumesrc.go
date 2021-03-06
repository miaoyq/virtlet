/*
Copyright 2018 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package libvirttools

// GetDefaultVolumeSource returns a volume source that supports
// root volume, flexvolumes and a ConfigSource for cloud-init
func GetDefaultVolumeSource() VMVolumeSource {
	return CombineVMVolumeSources(
		GetRootVolume,
		ScanFlexVolumes,
		// XXX: GetConfigVolume must go last because it
		// doesn't produce correct name for cdrom devices
		GetConfigVolume)
}
