/*

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

package alluxio

import (
	cruntime "github.com/fluid-cloudnative/fluid/pkg/runtime"
)

// SyncReplicas syncs the replicas
func (e *AlluxioEngine) SyncReplicas(ctx cruntime.ReconcileRequestContext) (err error) {

	runtime, err := e.getRuntime()
	if err != nil {
		return err
	}

	if runtime.Replicas() > runtime.Status.CurrentWorkerNumberScheduled {
		err = e.SetupWorkers()
		if err != nil {
			return err
		}
		_, err = e.CheckWorkersReady()
		if err != nil {
			e.Log.Error(err, "Check if the workers are ready")
			return err
		}

		// _, err := e.CheckAndUpdateRuntimeStatus()
		// if err != nil {
		// 	e.Log.Error(err, "Check if the runtime is ready")
		// 	return err
		// }
	} else if runtime.Replicas() < runtime.Status.CurrentWorkerNumberScheduled {
		// scale in
		e.Log.V(1).Info("Scale in to be implemented")
	} else {
		e.Log.V(1).Info("Nothing to do")
	}

	return

}
