# golang-restapi
<h2>Database Configuration</h2><br/>
<span>CREATE DATABASE gorestapidb;<br/>
USE gorestapidb;<br/>
CREATE TABLE songs (<br/>
    id INT AUTO_INCREMENT PRIMARY KEY,<br/>
    title VARCHAR(50) NOT NULL,<br/>
    singer VARCHAR(50) NOT NULL<br/>
);</span><br/>
<br/>
<h2> <strong>API Endpoint</strong></h2><br/>
<ul>
<li><a href="http://localhost:1234/api/v1/companies">http://localhost:1234/api/v1/companies</a>
<ul>
<li><code>GET</code>: get list of Users</li>
<li><code>POST</code>: create User</li>
</ul>
</li>
</ul>
<hr/>
<h2>Data Structure</h2><br/>
<div class="highlight highlight-source-json"><pre>{
  <span class="pl-s"><span class="pl-pds">"</span>name<span class="pl-pds">"</span></span>: <span class="pl-s"><span class="pl-pds">"</span>golang<span class="pl-pds">"</span></span>,
  <span class="pl-s"><span class="pl-pds">"</span>tel<span class="pl-pds">"</span></span>: <span class="pl-s"><span class="pl-pds">"</span>012-345-6789<span class="pl-pds">"</span></span>,
  <span class="pl-s"><span class="pl-pds">"</span>email<span class="pl-pds">"</span></span>: <span class="pl-s"><span class="pl-pds">"</span>golang-nuts@googlegroups.com<span class="pl-pds">"</span></span>
}</pre></div>
