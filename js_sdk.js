class ChurnSDK {
    constructor(apiKey) {
        if (!apiKey) {
            throw new Error("API key is required.");
        }
        this.apiKey = apiKey;
        this.apiUrl = "https://api.yourplatform.com/v1";
    }

    // Send data to the API
    async sendRequest(endpoint, data) {
        const url = this.apiUrl + endpoint;
        const response = await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${this.apiKey}`,
            },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error(`API request failed with status: ${response.status}`);
        }

        return response.json();
    }

    // Track user activity
    async trackUserActivity(userId, loginFrequency, lastLoginDate, featureUsage, sessionDuration, inAppActions) {
        const data = {
            user_id: userId,
            login_frequency: loginFrequency,
            last_login_date: lastLoginDate,
            feature_usage: featureUsage,
            session_duration: sessionDuration,
            in_app_actions: inAppActions,
        };
        return this.sendRequest("/user-activity", data);
    }

    // Track subscription data
    async trackSubscription(userId, subscriptionPlan, billingCycle, paymentStatus, subscriptionStartDate, subscriptionEndDate = null, cancellationDate = null) {
        const data = {
            user_id: userId,
            subscription_plan: subscriptionPlan,
            billing_cycle: billingCycle,
            payment_status: paymentStatus,
            subscription_start_date: subscriptionStartDate,
        };
        if (subscriptionEndDate) data.subscription_end_date = subscriptionEndDate;
        if (cancellationDate) data.cancellation_date = cancellationDate;
        return this.sendRequest("/subscription", data);
    }

    // Track engagement data
    async trackEngagement(userId, supportTickets, npsScore = null, csatScore = null, emailOpenRate = null, emailClickRate = null, feedback = null) {
        const data = {
            user_id: userId,
            support_tickets: supportTickets,
        };
        if (npsScore !== null) data.nps_score = npsScore;
        if (csatScore !== null) data.csat_score = csatScore;
        if (emailOpenRate !== null) data.email_open_rate = emailOpenRate;
        if (emailClickRate !== null) data.email_click_rate = emailClickRate;
        if (feedback !== null) data.feedback = feedback;
        return this.sendRequest("/engagement", data);
    }

    // Track custom metrics
    async trackCustomMetrics(userId, customMetrics) {
        const data = {
            user_id: userId,
            custom_metrics: customMetrics,
        };
        return this.sendRequest("/custom", data);
    }
}

// Export the SDK for use in Node.js or ES modules
if (typeof module !== "undefined" && module.exports) {
    module.exports = ChurnSDK; // For Node.js
} else {
    window.ChurnSDK = ChurnSDK; // For browser
}