package openid

import (
  "net/url"
  "strings"
)

func RedirectUrl(id, callbackUrl, realm string) (string, error) {
  return redirectUrl(id, callbackUrl, realm, urlGetter)
}

func redirectUrl(id, callbackUrl, realm string, getter httpGetter) (string, error) {
  opEndpoint, opLocalId, claimedId, err := discover(id, getter)
  if err != nil {
    return "", err
  }
  return buildRedirectUrl(opEndpoint, opLocalId, claimedId, callbackUrl, realm)
}

func buildRedirectUrl(opEndpoint, opLocalId, claimedId, returnTo, realm string) (string, error) {
  values := make(url.Values)
  values.Add("openid.ns", "http://specs.openid.net/auth/2.0")
  values.Add("openid.mode", "checkid_setup")
  values.Add("openid.return_to", returnTo)
  values.Add("openid.ns.sreg","http://openid.net/extensions/sreg/1.1")
  values.Add("openid.pape.preferred_auth_policies","")
  values.Add("openid.ns.pape","http://specs.openid.net/extensions/pape/1.0")
  values.Add("openid.ns","http://specs.openid.net/auth/2.0")
  values.Add("openid.sreg.required","nickname,fullname,email,timezone")
  values.Add("openid.lp.query_membership","_FAS_ALL_GROUPS_")
  values.Add("openid.ns.lp","http://ns.launchpad.net/2007/openid-teams")
  values.Add("openid.cla.query_cla","http://admin.fedoraproject.org/accounts/cla/done")
  values.Add("openid.ns.cla","http://fedoraproject.org/specs/open_id/cla")

  if len(claimedId) > 0 {
    values.Add("openid.claimed_id", claimedId)
    if len(opLocalId) > 0 {
      values.Add("openid.identity", opLocalId)
    } else {
      values.Add("openid.identity",
        "http://specs.openid.net/auth/2.0/identifier_select")
    }
  } else {
    values.Add("openid.identity",
      "http://specs.openid.net/auth/2.0/identifier_select")
  }

  if len(realm) > 0 {
    values.Add("openid.realm", realm)
  }

  if strings.Contains(opEndpoint, "?") {
    return opEndpoint + "&" + values.Encode(), nil
  }
  return opEndpoint + "?" + values.Encode(), nil
}
