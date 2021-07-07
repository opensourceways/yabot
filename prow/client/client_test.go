package client

import "testing"

func genToken() []byte {
	return []byte("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")
}

func TestClient_CreateIssue(t *testing.T) {
	client := NewClient(genToken)
	id, err := client.CreateIssue(
		"xwz123", "test",
		"yabot test create issue", "this just a test",
		1, []string{"kind/test"}, []string{})
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestClient_GetIssueLabels(t *testing.T) {
	client := NewClient(genToken)
	labels, err := client.GetIssueLabels("xwz123", "test", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(labels)
}
