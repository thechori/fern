using SeedApi.Core;

#nullable enable

namespace SeedApi.A.C;

public class CClient
{
    private RawClient _client;

    public CClient(RawClient client)
    {
        _client = client;
    }

    public async void FooAsync()
    {
        var response = await _client.MakeRequestAsync(
            new RawClient.JsonApiRequest { Method = HttpMethod.Post, Path = "" }
        );
    }
}
