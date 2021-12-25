const accountSid = 'ACc0efcb44a2df7fc58b302b0507592246'; 
const authToken = process.env['auth_token'];
const to_whatsapp_no = process.env['to_whatsapp_no'];
const from_whatsapp_no = process.env['from_whatsapp_no'];

const client = require('twilio')(accountSid, authToken); 
 
client.messages 
      .create({ 
         body: 'Your Yummy Cupcakes Company order of 1 dozen frosted cupcakes has shipped and should be delivered today. Details: http://www.baristobakery.com/', 
         from: 'whatsapp:+14155238886',       
         to: 'whatsapp:'+to_whatsapp_no
 
       }) 
      .then(message => console.log(message.sid)) 
      .done();