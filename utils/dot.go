package utils

import (
    "bytes"
    "fmt"
    "strings"
    "os/exec"
    "tree-visualization/models"
)

func GenerateDOT(nodes []models.Node) string {
    var builder strings.Builder
    builder.WriteString("digraph G {\n")
    
    // Edges
    for _, node := range nodes {
        if node.ParentID != nil {
            builder.WriteString(fmt.Sprintf("    \"%d\" -> \"%d\";\n", 
                *node.ParentID, node.ID))
        }
    }
    
    // Nodes with labels
    for _, node := range nodes {
        builder.WriteString(fmt.Sprintf("    \"%d\" [label=\"%s\\n(left: %d, right: %d)\"];\n",
            node.ID, node.DisplayName, node.NodeLeft, node.NodeRight))
    }
    
    builder.WriteString("}\n")
    return builder.String()
}

func GenerateImage(dotContent string, format string) ([]byte, error) {
    cmd := exec.Command("dot", "-T"+format)
    cmd.Stdin = strings.NewReader(dotContent)
    var out bytes.Buffer
    cmd.Stdout = &out
    
    err := cmd.Run()
    if err != nil {
        return nil, fmt.Errorf("error generating image: %v", err)
    }
    return out.Bytes(), nil
}