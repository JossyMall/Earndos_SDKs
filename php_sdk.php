<?php
class ChurnSDK {
    private $apiKey;
    private $apiUrl = 'https://api.yourplatform.com/v1';

    public function __construct($apiKey) {
        if (empty($apiKey)) {
            throw new Exception("API key is required.");
        }
        $this->apiKey = $apiKey;
    }

    // Send data to the API
    private function sendRequest($endpoint, $data) {
        $url = $this->apiUrl . $endpoint;
        $options = [
            'http' => [
                'header'  => "Content-type: application/json\r\nAuthorization: Bearer " . $this->apiKey,
                'method'  => 'POST',
                'content' => json_encode($data),
            ],
        ];
        $context = stream_context_create($options);
        $response = file_get_contents($url, false, $context);

        if ($response === FALSE) {
            throw new Exception("Failed to send data to the API.");
        }

        return json_decode($response, true);
    }

    // Track user activity
    public function trackUserActivity($userId, $loginFrequency, $lastLoginDate, $featureUsage, $sessionDuration, $inAppActions) {
        $endpoint = '/user-activity';
        $data = [
            'user_id' => $userId,
            'login_frequency' => $loginFrequency,
            'last_login_date' => $lastLoginDate,
            'feature_usage' => $featureUsage,
            'session_duration' => $sessionDuration,
            'in_app_actions' => $inAppActions,
        ];
        return $this->sendRequest($endpoint, $data);
    }

    // Track subscription data
    public function trackSubscription($userId, $subscriptionPlan, $billingCycle, $paymentStatus, $subscriptionStartDate, $subscriptionEndDate = null, $cancellationDate = null) {
        $endpoint = '/subscription';
        $data = [
            'user_id' => $userId,
            'subscription_plan' => $subscriptionPlan,
            'billing_cycle' => $billingCycle,
            'payment_status' => $paymentStatus,
            'subscription_start_date' => $subscriptionStartDate,
            'subscription_end_date' => $subscriptionEndDate,
            'cancellation_date' => $cancellationDate,
        ];
        return $this->sendRequest($endpoint, $data);
    }

    // Track engagement data
    public function trackEngagement($userId, $supportTickets, $npsScore = null, $csatScore = null, $emailOpenRate = null, $emailClickRate = null, $feedback = null) {
        $endpoint = '/engagement';
        $data = [
            'user_id' => $userId,
            'support_tickets' => $supportTickets,
            'nps_score' => $npsScore,
            'csat_score' => $csatScore,
            'email_open_rate' => $emailOpenRate,
            'email_click_rate' => $emailClickRate,
            'feedback' => $feedback,
        ];
        return $this->sendRequest($endpoint, $data);
    }

    // Track custom metrics
    public function trackCustomMetrics($userId, $customMetrics) {
        $endpoint = '/custom';
        $data = [
            'user_id' => $userId,
            'custom_metrics' => $customMetrics,
        ];
        return $this->sendRequest($endpoint, $data);
    }
}
?>