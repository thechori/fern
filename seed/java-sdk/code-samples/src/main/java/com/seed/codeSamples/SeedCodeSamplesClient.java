/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.codeSamples;

import com.seed.codeSamples.core.ClientOptions;
import com.seed.codeSamples.core.Suppliers;
import com.seed.codeSamples.resources.service.ServiceClient;
import java.util.function.Supplier;

public class SeedCodeSamplesClient {
    protected final ClientOptions clientOptions;

    protected final Supplier<ServiceClient> serviceClient;

    public SeedCodeSamplesClient(ClientOptions clientOptions) {
        this.clientOptions = clientOptions;
        this.serviceClient = Suppliers.memoize(() -> new ServiceClient(clientOptions));
    }

    public ServiceClient service() {
        return this.serviceClient.get();
    }

    public static SeedCodeSamplesClientBuilder builder() {
        return new SeedCodeSamplesClientBuilder();
    }
}