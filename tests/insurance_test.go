package easypost_test

import (
	"reflect"
	"strings"

	"github.com/EasyPost/easypost-go/v2"
)

func (c *ClientTests) TestInsuranceCreate() {
	client := c.TestClient()
	assert, require := c.Assert(), c.Require()

	shipment, err := client.CreateShipment(c.fixture.OneCallBuyShipment())
	require.NoError(err)

	insuranceData := c.fixture.BasicInsurance()
	insuranceData.TrackingCode = shipment.TrackingCode

	insurance, err := client.CreateInsurance(insuranceData)
	require.NoError(err)

	assert.Equal(reflect.TypeOf(&easypost.Insurance{}), reflect.TypeOf(insurance))
	assert.True(strings.HasPrefix(insurance.ID, "ins_"))
	assert.Equal("100.00000", insurance.Amount)
}

func (c *ClientTests) TestInsuranceRetrieve() {
	client := c.TestClient()
	assert, require := c.Assert(), c.Require()

	shipment, err := client.CreateShipment(c.fixture.OneCallBuyShipment())
	require.NoError(err)

	insuranceData := c.fixture.BasicInsurance()
	insuranceData.TrackingCode = shipment.TrackingCode

	insurance, err := client.CreateInsurance(insuranceData)
	require.NoError(err)

	retrievedInsurance, err := client.GetInsurance(insurance.ID)
	require.NoError(err)

	assert.Equal(reflect.TypeOf(&easypost.Insurance{}), reflect.TypeOf(retrievedInsurance))
	assert.Equal(insurance, retrievedInsurance)
}

func (c *ClientTests) TestInsuranceAll() {
	client := c.TestClient()
	assert, require := c.Assert(), c.Require()

	insurances, err := client.ListInsurances(
		&easypost.ListOptions{
			PageSize: c.fixture.pageSize(),
		},
	)
	require.NoError(err)

	insurancesList := insurances.Insurances

	assert.LessOrEqual(len(insurancesList), c.fixture.pageSize())
	assert.NotNil(insurances.HasMore)
	for _, insurance := range insurancesList {
		assert.Equal(reflect.TypeOf(&easypost.Insurance{}), reflect.TypeOf(insurance))
	}
}
