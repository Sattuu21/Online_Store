import React, { useState, useEffect } from "react";
import { Elements } from "@stripe/react-stripe-js";
import { loadStripe } from "@stripe/stripe-js";
import PaymentForm from "./PaymentForm"; // Ensure this path is correct

const stripePromise = loadStripe("pk_test_51PehjqRsBzWtaf8b3OTB0c2JAcivMWtr0ncB8T8ZlxIaZbBmhZV2Lpud6twoIZgdRTDJ4wUcomBmBEVnKDxKxNlG00LhbPXoRZ"); 

export default function StripePayment({ customerData }) {
  const [clientSecret, setClientSecret] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    const createPaymentIntent = async () => {
      try {
        const response = await fetch("/create-payment-intent", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(customerData),
        });

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();

        if (data.clientSecret) {
          setClientSecret(data.clientSecret);
        } else {
          throw new Error("No client secret in response");
        }
      } catch (error) {
        console.error("Error creating PaymentIntent:", error);
        setError(error.message);
      }
    };

   createPaymentIntent();
  }, [customerData]);

  const appearance = {
    theme: "stripe",
  };
  const options = {
    clientSecret,
    appearance,
  };

  return (
    <div className="StripePayment" style={{ height: "100%", width: "100%" }}>
      {error && <div>Error: {error}</div>}
      {clientSecret ? (
        <Elements options={options} stripe={stripePromise}>
          <PaymentForm />
        </Elements>
      ) : (
        !error && <div>Loading payment details...</div>
      )}
    </div>
  );
}
