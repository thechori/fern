/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../../serialization/index";
import * as SeedExamples from "../../../index";
import express from "express";
import * as errors from "../../../../errors/index";

export interface ServiceServiceMethods {
    getMovie(
        req: express.Request<
            {
                movieId: serializers.MovieId.Raw;
            },
            SeedExamples.Movie,
            never,
            never
        >,
        res: {
            send: (responseBody: SeedExamples.Movie) => Promise<void>;
            cookie: (cookie: string, value: string, options?: express.CookieOptions) => void;
            locals: any;
        }
    ): void | Promise<void>;
    createMovie(
        req: express.Request<never, SeedExamples.MovieId, SeedExamples.Movie, never>,
        res: {
            send: (responseBody: SeedExamples.MovieId) => Promise<void>;
            cookie: (cookie: string, value: string, options?: express.CookieOptions) => void;
            locals: any;
        }
    ): void | Promise<void>;
    getMetadata(
        req: express.Request<
            never,
            SeedExamples.Metadata,
            never,
            {
                shallow?: boolean;
                tag?: string;
            }
        >,
        res: {
            send: (responseBody: SeedExamples.Metadata) => Promise<void>;
            cookie: (cookie: string, value: string, options?: express.CookieOptions) => void;
            locals: any;
        }
    ): void | Promise<void>;
    getResponse(
        req: express.Request<never, SeedExamples.Response, never, never>,
        res: {
            send: (responseBody: SeedExamples.Response) => Promise<void>;
            cookie: (cookie: string, value: string, options?: express.CookieOptions) => void;
            locals: any;
        }
    ): void | Promise<void>;
}

export class ServiceService {
    private router;

    constructor(private readonly methods: ServiceServiceMethods, middleware: express.RequestHandler[] = []) {
        this.router = express.Router({ mergeParams: true }).use(
            express.json({
                strict: false,
            }),
            ...middleware
        );
    }

    public addMiddleware(handler: express.RequestHandler): this {
        this.router.use(handler);
        return this;
    }

    public toRouter(): express.Router {
        this.router.get("/movie/:movieId", async (req, res, next) => {
            try {
                await this.methods.getMovie(req as any, {
                    send: async (responseBody) => {
                        res.json(
                            await serializers.Movie.jsonOrThrow(responseBody, { unrecognizedObjectKeys: "strip" })
                        );
                    },
                    cookie: res.cookie.bind(res),
                    locals: res.locals,
                });
                next();
            } catch (error) {
                if (error instanceof errors.SeedExamplesError) {
                    console.warn(
                        `Endpoint 'getMovie' unexpectedly threw ${error.constructor.name}.` +
                            ` If this was intentional, please add ${error.constructor.name} to` +
                            " the endpoint's errors list in your Fern Definition."
                    );
                    await error.send(res);
                } else {
                    res.status(500).json("Internal Server Error");
                }
                next(error);
            }
        });
        this.router.post("/movie", async (req, res, next) => {
            const request = await serializers.Movie.parse(req.body);
            if (request.ok) {
                req.body = request.value;
                try {
                    await this.methods.createMovie(req as any, {
                        send: async (responseBody) => {
                            res.json(
                                await serializers.MovieId.jsonOrThrow(responseBody, { unrecognizedObjectKeys: "strip" })
                            );
                        },
                        cookie: res.cookie.bind(res),
                        locals: res.locals,
                    });
                    next();
                } catch (error) {
                    if (error instanceof errors.SeedExamplesError) {
                        console.warn(
                            `Endpoint 'createMovie' unexpectedly threw ${error.constructor.name}.` +
                                ` If this was intentional, please add ${error.constructor.name} to` +
                                " the endpoint's errors list in your Fern Definition."
                        );
                        await error.send(res);
                    } else {
                        res.status(500).json("Internal Server Error");
                    }
                    next(error);
                }
            } else {
                res.status(422).json({
                    errors: request.errors.map(
                        (error) => ["request", ...error.path].join(" -> ") + ": " + error.message
                    ),
                });
                next(request.errors);
            }
        });
        this.router.get("/metadata", async (req, res, next) => {
            try {
                await this.methods.getMetadata(req as any, {
                    send: async (responseBody) => {
                        res.json(
                            await serializers.Metadata.jsonOrThrow(responseBody, { unrecognizedObjectKeys: "strip" })
                        );
                    },
                    cookie: res.cookie.bind(res),
                    locals: res.locals,
                });
                next();
            } catch (error) {
                if (error instanceof errors.SeedExamplesError) {
                    console.warn(
                        `Endpoint 'getMetadata' unexpectedly threw ${error.constructor.name}.` +
                            ` If this was intentional, please add ${error.constructor.name} to` +
                            " the endpoint's errors list in your Fern Definition."
                    );
                    await error.send(res);
                } else {
                    res.status(500).json("Internal Server Error");
                }
                next(error);
            }
        });
        this.router.post("/response", async (req, res, next) => {
            try {
                await this.methods.getResponse(req as any, {
                    send: async (responseBody) => {
                        res.json(
                            await serializers.Response.jsonOrThrow(responseBody, { unrecognizedObjectKeys: "strip" })
                        );
                    },
                    cookie: res.cookie.bind(res),
                    locals: res.locals,
                });
                next();
            } catch (error) {
                if (error instanceof errors.SeedExamplesError) {
                    console.warn(
                        `Endpoint 'getResponse' unexpectedly threw ${error.constructor.name}.` +
                            ` If this was intentional, please add ${error.constructor.name} to` +
                            " the endpoint's errors list in your Fern Definition."
                    );
                    await error.send(res);
                } else {
                    res.status(500).json("Internal Server Error");
                }
                next(error);
            }
        });
        return this.router;
    }
}