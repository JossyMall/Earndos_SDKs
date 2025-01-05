<?php
session_start();
require_once '../db.php'; // Include the database connection

// Check if the user is logged in
if (!isset($_SESSION['user_id'])) {
    header('Location: ../auth/login.php');
    exit();
}

$user_id = $_SESSION['user_id'];

// Fetch the user's active API key
$stmt = $pdo->prepare("SELECT api_key FROM api_keys WHERE user_id = ? AND is_active = 1");
$stmt->execute([$user_id]);
$api_key = $stmt->fetchColumn();
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PHP SDK Documentation</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="container">
        <h1>PHP SDK Documentation</h1>

        <?php if (!$api_key): ?>
            <div class="message">You need to create an API key to use the SDK. <a href="api_keys.php">Create API Key</a></div>
        <?php else: ?>
            <h2>Installation</h2>
            <p>Download the PHP SDK and include it in your project:</p>
            <pre><code>&lt;?php
require_once 'path/to/php_sdk.php';
?&gt;</code></pre>

            <h2>Initialization</h2>
            <p>Initialize the SDK with your API key:</p>
            <pre><code>&lt;?php
$churnSDK = new ChurnSDK('<?= $api_key ?>');
?&gt;</code></pre>

            <h2>Tracking User Activity</h2>
            <p>Example: Track user login frequency, feature usage, and session duration:</p>
            <pre><code>&lt;?php
$churnSDK->trackUserActivity('12345', 10, '2023-10-01', ['dashboard' => 15, 'reports' => 5], 300, ['clicked_button', 'downloaded_file']);
?&gt;</code></pre>

            <h2>Tracking Subscription Data</h2>
            <p>Example: Track subscription plan, billing cycle, and payment status:</p>
            <pre><code>&lt;?php
$churnSDK->trackSubscription('12345', 'premium', 'monthly', 'success', '2023-01-01');
?&gt;</code></pre>

            <h2>Tracking Engagement Data</h2>
            <p>Example: Track support tickets, NPS score, and email engagement:</p>
            <pre><code>&lt;?php
$churnSDK->trackEngagement('12345', 2, 9, 4.5, 0.75, 0.25, 'The platform is great!');
?&gt;</code></pre>

            <h2>Tracking Custom Metrics</h2>
            <p>Example: Track custom metrics specific to your application:</p>
            <pre><code>&lt;?php
$churnSDK->trackCustomMetrics('12345', ['projects_created' => 10, 'tasks_completed' => 50]);
?&gt;</code></pre>

            <h2>Error Handling</h2>
            <p>Always wrap SDK calls in a try-catch block to handle errors gracefully:</p>
            <pre><code>&lt;?php
try {
    $churnSDK->trackUserActivity('12345', 10, '2023-10-01', ['dashboard' => 15, 'reports' => 5], 300, ['clicked_button', 'downloaded_file']);
} catch (Exception $e) {
    echo "Error: " . $e->getMessage();
}
?&gt;</code></pre>
        <?php endif; ?>

        <p><a href="dashboard.php">Back to Dashboard</a></p>
    </div>
</body>
</html>