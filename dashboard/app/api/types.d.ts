/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/auth/login": {
    /**
     * Session Token Authentication.
     * @description Login to the service and retrieve a session token for authentication.
     */
    post: operations["post-auth-login"];
  };
  "/event/hit": {
    /**
     * Send Hit Event.
     * @description Send a hit event to register a user view.
     */
    post: operations["post-event-hit"];
  };
  "/event/ping": {
    /**
     * Unique User Check.
     * @description This is a ping endpoint to determine if the user is unique or not.
     */
    get: operations["get-event-ping"];
  };
  "/user": {
    /**
     * Get User Info.
     * @description Retrieve the information of the user with the matching user ID.
     */
    get: operations["get-user"];
    /**
     * Create New User.
     * @description Create a new user.
     */
    post: operations["post-user"];
    /**
     * Delete User.
     * @description Delete a user account.
     */
    delete: operations["delete-user"];
    /**
     * Update User Info.
     * @description Update a user account's details.
     */
    patch: operations["patch-user"];
  };
  "/websites": {
    /**
     * List Websites.
     * @description Get a list of all websites from the user.
     */
    get: operations["get-websites"];
    /**
     * Add Website.
     * @description Add a new website.
     */
    post: operations["post-websites"];
  };
  "/websites/{hostname}": {
    /**
     * Get Website.
     * @description Get website details for an individual website.
     */
    get: operations["get-websites-id"];
    /**
     * Delete Website.
     * @description Delete a website.
     */
    delete: operations["delete-websites-id"];
    /**
     * Update Website.
     * @description Update a website's information.
     */
    patch: operations["patch-websites-id"];
  };
  "/website/{hostname}/summary": {
    /**
     * Get Stat Summary.
     * @description Get a summary of the website's stats.
     */
    get: operations["get-website-id-summary"];
  };
  "/website/{hostname}/pages": {
    /**
     * Get Page Stats.
     * @description Get a list of pages and their stats.
     */
    get: operations["get-website-id-pages"];
  };
  "/website/{hostname}/time": {
    /**
     * Get Time Stats.
     * @description Get a list of pages and their time stats.
     */
    get: operations["get-website-id-time"];
  };
  "/website/{hostname}/referrers": {
    /**
     * Get Referrer Stats.
     * @description Get a list of referrers and their stats.
     */
    get: operations["get-website-id-referrers"];
  };
  "/website/{hostname}/sources": {
    /**
     * Get UTM Source Stats.
     * @description Get a list of UTM sources and their stats.
     */
    get: operations["get-website-id-sources"];
  };
  "/website/{hostname}/mediums": {
    /**
     * Get UTM Medium Stats.
     * @description Get a list of UTM mediums and their stats.
     */
    get: operations["get-website-id-mediums"];
  };
  "/website/{hostname}/campaigns": {
    /**
     * Get UTM Campaign Stats.
     * @description Get a list of UTM campaigns and their stats.
     */
    get: operations["get-website-id-campaigns"];
  };
  "/website/{hostname}/browsers": {
    /**
     * Get Browser Stats.
     * @description Get a list of browsers and their stats.
     */
    get: operations["get-website-id-browsers"];
  };
  "/website/{hostname}/os": {
    /**
     * Get OS Stats.
     * @description Get a list of OS and their stats.
     */
    get: operations["get-website-id-os"];
  };
  "/website/{hostname}/devices": {
    /**
     * Get Device Stats.
     * @description Get a list of devices and their stats.
     */
    get: operations["get-website-id-device"];
  };
  "/website/{hostname}/screens": {
    /**
     * Get Screen Stats.
     * @description Get a list of screen sizes and their stats.
     */
    get: operations["get-website-id-screen"];
  };
  "/website/{hostname}/countries": {
    /**
     * Get Country Stats.
     * @description Get a list of countries and their stats.
     */
    get: operations["get-website-id-country"];
  };
  "/website/{hostname}/languages": {
    /**
     * Get Language Stats.
     * @description Get a list of languages and their stats.
     */
    get: operations["get-website-id-language"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    /**
     * AuthLogin
     * @description Request body for logging in.
     */
    AuthLogin: {
      /** Format: email */
      email: string;
      /** Format: password */
      password: string;
    };
    /**
     * EventHit
     * @description Website hit event.
     */
    EventHit: {
      /** @description UUID generated for each user to link multiple events on the same page together. */
      b: string;
      /** @description Page URL including query parameters. */
      u: string;
      /** @description Referrer URL. */
      r?: string;
      /** @description If the user is a unique user or not. */
      p?: boolean;
      /** @description Event type consisting of either 'pagehide', 'unload', 'load', 'hidden' or 'visible'. */
      e: string;
      /** @description Title of page. */
      t?: string;
      /** @description Timezone of the user. */
      d?: string;
      /** @description Screen width. */
      w?: number;
      /** @description Screen height. */
      h?: number;
      /** @description Time spent on page. Only sent on unload. */
      m?: number;
    };
    /**
     * UserCreate
     * @description Request body for creating a user.
     */
    UserCreate: {
      /** Format: email */
      email: string;
      /** Format: password */
      password: string;
      /**
       * @default en
       * @enum {string}
       */
      language?: "en";
    };
    /**
     * UserGet
     * @description Response body for getting a user.
     */
    UserGet: {
      /** Format: email */
      email: string;
      /**
       * @default en
       * @enum {string}
       */
      language: "en";
      /** Format: int64 */
      dateCreated: number;
      /** Format: int64 */
      dateUpdated: number;
    };
    /**
     * UserPatch
     * @description Request body for updating a user.
     */
    UserPatch: {
      /** Format: email */
      email?: string;
      /** Format: password */
      password?: string;
      /**
       * @default en
       * @enum {string}
       */
      language?: "en";
    };
    /**
     * WebsiteGet
     * @description Response body for getting a website.
     */
    WebsiteGet: {
      name: string;
      /** Format: hostname */
      hostname: string;
    };
    /**
     * WebsiteCreate
     * @description Request body for creating a website.
     */
    WebsiteCreate: {
      name: string;
      /** Format: hostname */
      hostname: string;
    };
    /**
     * WebsitePatch
     * @description Request body for updating a website.
     */
    WebsitePatch: {
      name?: string;
      /** Format: hostname */
      hostname?: string;
    };
    /** StatsSummary */
    StatsSummary: {
      uniques: number;
      pageviews: number;
      bounces: number;
      duration: number;
      active?: number;
    };
    /** StatsPages */
    StatsPages: {
        /** @description Pathname of the page. */
        path: string;
        /** @description Title of the page. */
        title?: string;
        /** @description Number of unique users. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users.
         */
        unique_percentage: number;
        /** @description Number of page views. */
        pageviews?: number;
        /** @description Number of bounces. */
        bounces?: number;
        /** @description Total time spent on page in milliseconds. */
        duration?: number;
      }[];
    /** StatsTime */
    StatsTime: {
        /** @description Pathname of the page. */
        path: string;
        /** @description Title of the page. */
        title?: string;
        /** @description Median time spent on page in milliseconds. */
        duration: number;
        /** @description Total time spent on page in milliseconds for the upper quartile (75%). */
        duration_upper_quartile?: number;
        /** @description Total time spent on page in milliseconds for the lower quartile (25%). */
        duration_lower_quartile?: number;
        /**
         * Format: float
         * @description Percentage of time contributing to the total time spent on the website.
         */
        duration_percentage: number;
        /** @description Number of unique users. */
        uniques?: number;
        /** @description Number of bounces. */
        bounces?: number;
      }[];
    /** StatsReferrers */
    StatsReferrers: {
        /** @description Referrer URL. */
        referrer: string;
        /** @description Number of unique users from referrer. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from referrer.
         */
        unique_percentage: number;
        /** @description Number of bounces from referrer. */
        bounces?: number;
        /** @description Total time spent on page from referrer in milliseconds. */
        duration?: number;
      }[];
    /** StatsUTMSources */
    StatsUTMSources: {
        /** @description UTM source. */
        source: string;
        /** @description Number of unique users from UTM source. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from UTM source.
         */
        uniquepercentage: number;
      }[];
    /** StatsUTMMediums */
    StatsUTMMediums: {
        /** @description UTM medium. */
        medium: string;
        /** @description Number of unique users from UTM medium. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from UTM medium.
         */
        uniquepercentage: number;
      }[];
    /** StatsUTMCampaigns */
    StatsUTMCampaigns: {
        /** @description UTM campaign. */
        campaign: string;
        /** @description Number of unique users from UTM campaign. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from UTM campaign.
         */
        uniquepercentage: number;
      }[];
    /** StatsBrowsers */
    StatsBrowsers: {
        /** @description Browser name. */
        browser: string;
        /** @description Number of unique users from browser. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from browser.
         */
        uniquepercentage: number;
        /** @description Browser version. */
        version?: string;
      }[];
    /** StatsOS */
    StatsOS: {
        /** @description OS name. */
        os: string;
        /** @description Number of unique users from OS. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from OS.
         */
        uniquepercentage: number;
        /** @description OS version. */
        version?: string;
      }[];
    /** StatsDevices */
    StatsDevices: {
        /** @description Device name. */
        device: string;
        /** @description Number of unique users from device. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from device.
         */
        uniquepercentage: number;
      }[];
    /** StatsScreens */
    StatsScreens: {
        /** @description Screen size. */
        screen: string;
        /** @description Number of unique users from screen size. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from screen size.
         */
        uniquepercentage: number;
      }[];
    /** StatsCountries */
    StatsCountries: {
        /** @description Country name. */
        country: string;
        /** @description Number of unique users from country. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from country.
         */
        uniquepercentage: number;
      }[];
    /** StatsLanguages */
    StatsLanguages: {
        /** @description Language name. */
        language: string;
        /** @description Number of unique users from language. */
        uniques: number;
        /**
         * Format: float
         * @description Percentage of unique users from language.
         */
        uniquepercentage: number;
      }[];
  };
  responses: {
    /** @description 400 Bad Request. */
    BadRequestError: {
      content: {
        "application/json": {
          error: {
            /**
             * Format: int32
             * @default 400
             */
            code: number;
            message: string;
          };
        };
      };
    };
    /** @description 401 Unauthorised. */
    UnauthorisedError: {
      content: {
        "application/json": {
          error: {
            /**
             * Format: int32
             * @default 401
             */
            code: number;
            message: string;
          };
        };
      };
    };
    /** @description 403 Forbidden. */
    ForbiddenError: {
      content: {
        "application/json": {
          error: {
            /**
             * Format: int32
             * @default 403
             */
            code: number;
            message: string;
          };
        };
      };
    };
    /** @description 404 Not Found. */
    NotFoundError: {
      content: {
        "application/json": {
          error: {
            /**
             * Format: int32
             * @default 404
             */
            code: number;
            message: string;
          };
        };
      };
    };
    /** @description 409 Conflict Found. */
    ConflictError: {
      content: {
        "application/json": {
          error: {
            /**
             * Format: int32
             * @default 409
             */
            code: number;
            message: string;
          };
        };
      };
    };
    /** @description 500 Unexpected Internal Server Error. */
    InternalServerError: {
      content: {
        "application/json": {
          error: {
            /**
             * Format: int32
             * @default 500
             */
            code: number;
            message: string;
          };
        };
      };
    };
  };
  parameters: {
    /** @description Session token for authentication. */
    SessionAuth: string;
    /** @description Hostname for the website. */
    Hostname: string;
    /** @description Return a summary of the stats. */
    Summary?: boolean;
    /** @description Start time (seconds) in Unix epoch format. */
    PeriodStart?: string;
    /** @description End time (seconds) in Unix epoch format. */
    PeriodEnd?: string;
    /** @description Limit the number of results. */
    Limit?: number;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /**
   * Session Token Authentication.
   * @description Login to the service and retrieve a session token for authentication.
   */
  "post-auth-login": {
    /** @description Login details. */
    requestBody: {
      content: {
        "application/json": components["schemas"]["AuthLogin"];
      };
    };
    responses: {
      /** @description Success */
      200: {
        headers: {
          /** @description Set the cookie for the session. */
          "Set-Cookie": string;
        };
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Send Hit Event.
   * @description Send a hit event to register a user view.
   */
  "post-event-hit": {
    parameters: {
      header?: {
        /** @description Used to infer user browser, OS and device. */
        "User-Agent"?: string;
        /** @description Used to infer user language. */
        "Accept-Language"?: string;
      };
    };
    /** @description Hit event metadata. */
    requestBody: {
      content: {
        "application/json": components["schemas"]["EventHit"];
      };
    };
    responses: {
      /** @description OK. */
      200: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Unique User Check.
   * @description This is a ping endpoint to determine if the user is unique or not.
   */
  "get-event-ping": {
    parameters: {
      header?: {
        /** @description If this exists, then user exists in cache and is not a unique user. */
        "If-Modified-Since"?: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        headers: {
          /** @description This is date of the current day from midnight, incremented by each page view by unique user. */
          "Last-Modified": string;
          /** @description This is set to 1 day to prevent the user from being counted as a unique user again. */
          "Cache-Control": string;
          /** @description This is set to allow all origins to access this endpoint. */
          "Access-Control-Allow-Origin": string;
        };
        content: {
          "text/plain": string;
        };
      };
      400: components["responses"]["BadRequestError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get User Info.
   * @description Retrieve the information of the user with the matching user ID.
   */
  "get-user": {
    parameters: {
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description User Found */
      200: {
        content: {
          "application/json": components["schemas"]["UserGet"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Create New User.
   * @description Create a new user.
   */
  "post-user": {
    /** @description Post the necessary fields for the API to create a new user. */
    requestBody: {
      content: {
        "application/json": components["schemas"]["UserCreate"];
      };
    };
    responses: {
      /** @description User Created */
      201: {
        headers: {
          /** @description Set the cookie for the session. */
          "Set-Cookie": string;
        };
        content: {
          "application/json": components["schemas"]["UserGet"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Delete User.
   * @description Delete a user account.
   */
  "delete-user": {
    parameters: {
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description Success */
      200: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update User Info.
   * @description Update a user account's details.
   */
  "patch-user": {
    parameters: {
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    /** @description User details to update. */
    requestBody: {
      content: {
        "application/json": components["schemas"]["UserPatch"];
      };
    };
    responses: {
      /** @description Success */
      200: {
        content: {
          "application/json": components["schemas"]["UserGet"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * List Websites.
   * @description Get a list of all websites from the user.
   */
  "get-websites": {
    parameters: {
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description Returns a list of websites. */
      200: {
        content: {
          "application/json": components["schemas"]["WebsiteGet"][];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Add Website.
   * @description Add a new website.
   */
  "post-websites": {
    requestBody: {
      content: {
        "application/json": components["schemas"]["WebsiteCreate"];
      };
    };
    responses: {
      /** @description Created */
      201: {
        content: {
          "application/json": components["schemas"]["WebsiteGet"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Website.
   * @description Get website details for an individual website.
   */
  "get-websites-id": {
    parameters: {
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["WebsiteGet"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Delete Website.
   * @description Delete a website.
   */
  "delete-websites-id": {
    parameters: {
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description Success */
      200: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update Website.
   * @description Update a website's information.
   */
  "patch-websites-id": {
    parameters: {
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    /** @description Website details to update. */
    requestBody: {
      content: {
        "application/json": components["schemas"]["WebsitePatch"];
      };
    };
    responses: {
      /** @description Success */
      200: {
        content: {
          "application/json": components["schemas"]["WebsiteGet"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Stat Summary.
   * @description Get a summary of the website's stats.
   */
  "get-website-id-summary": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsSummary"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Page Stats.
   * @description Get a list of pages and their stats.
   */
  "get-website-id-pages": {
    parameters: {
      query?: {
        summary?: components["parameters"]["Summary"];
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsPages"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Time Stats.
   * @description Get a list of pages and their time stats.
   */
  "get-website-id-time": {
    parameters: {
      query?: {
        summary?: components["parameters"]["Summary"];
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsTime"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Referrer Stats.
   * @description Get a list of referrers and their stats.
   */
  "get-website-id-referrers": {
    parameters: {
      query?: {
        summary?: components["parameters"]["Summary"];
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsReferrers"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get UTM Source Stats.
   * @description Get a list of UTM sources and their stats.
   */
  "get-website-id-sources": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsUTMSources"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get UTM Medium Stats.
   * @description Get a list of UTM mediums and their stats.
   */
  "get-website-id-mediums": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsUTMMediums"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get UTM Campaign Stats.
   * @description Get a list of UTM campaigns and their stats.
   */
  "get-website-id-campaigns": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsUTMCampaigns"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Browser Stats.
   * @description Get a list of browsers and their stats.
   */
  "get-website-id-browsers": {
    parameters: {
      query?: {
        summary?: components["parameters"]["Summary"];
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsBrowsers"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get OS Stats.
   * @description Get a list of OS and their stats.
   */
  "get-website-id-os": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsOS"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Device Stats.
   * @description Get a list of devices and their stats.
   */
  "get-website-id-device": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsDevices"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Screen Stats.
   * @description Get a list of screen sizes and their stats.
   */
  "get-website-id-screen": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsScreens"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Country Stats.
   * @description Get a list of countries and their stats.
   */
  "get-website-id-country": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsCountries"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Language Stats.
   * @description Get a list of languages and their stats.
   */
  "get-website-id-language": {
    parameters: {
      query?: {
        start?: components["parameters"]["PeriodStart"];
        end?: components["parameters"]["PeriodEnd"];
        limit?: components["parameters"]["Limit"];
      };
      path: {
        hostname: components["parameters"]["Hostname"];
      };
      cookie: {
        _me_sess: components["parameters"]["SessionAuth"];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["StatsLanguages"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorisedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
}
