/*
Copyright The KubeDB Authors.

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
package framework

import (
	"time"

	. "github.com/onsi/gomega"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyAppBinding(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.appCatalogClient.AppBindings(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			if err != nil {
				if kerr.IsNotFound(err) {
					return false
				}
				Expect(err).NotTo(HaveOccurred())
			}
			return true
		},
		time.Minute*5,
		time.Second*5,
	)
}

func (f *Framework) CheckPostgresAppBindingSpec(meta metav1.ObjectMeta) error {
	postgres, err := f.GetPostgres(meta)
	Expect(err).NotTo(HaveOccurred())

	_, err = f.appCatalogClient.AppBindings(postgres.Namespace).Get(postgres.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (f *Framework) CheckPgBouncerAppBindingSpec(meta metav1.ObjectMeta) error {
	pgbouncer, err := f.GetPgBouncer(meta)
	Expect(err).NotTo(HaveOccurred())

	_, err = f.appCatalogClient.AppBindings(pgbouncer.Namespace).Get(pgbouncer.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	return nil
}