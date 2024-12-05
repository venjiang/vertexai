# vertexai

## IAM 角色

如需使用 Vertex AI 上的生成式 AI 功能，需要为您项目中的主账号（例如用户、群组和服务账号）授予适当的 IAM：角色。您还可以创建自定义角色，以向主账号授予一组用户定义的权限。

**预定义角色**
您可以向项目中的用户或群组授予以下预定义角色之一，以授予他们访问 Vertex AI 上的生成式 AI 功能的权限：

- Vertex AI Administrator (roles/aiplatform.admin)

- Vertex AI User (roles/aiplatform.user)

如需详细了解 Vertex AI IAM 角色，请参阅:

- [使用 IAM 进行 Vertex AI 访问权限控制](https://cloud.google.com/vertex-ai/docs/general/access-control?hl=zh-cn)。

- [支持进行服务账号关联的 Google Cloud 服务](https://cloud.google.com/docs/authentication/provide-credentials-adc?hl=zh-cn#attached-sa)
