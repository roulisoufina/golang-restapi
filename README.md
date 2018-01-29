# golang-restapi
Simple Restful API using Golang and MYSQL<br/>
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
    <li> // create new song<br/>
        POST: <a href="#">http://localhost:5555/song</a><br/>
        body:<br/>
        {<br/>
	    "Title":"Fight Song",<br/>
	    "Singer":"Rachel Platten"<br/>
        }<br/>
    </li>
    <li> // get all songs<br/>
        GET: <a href="#">http://localhost:5555/songs</a><br/>
    </li>
    <li> // get specific song<br/>
        GET: <a href="#">http://localhost:5555/song/3</a><br/>
    </li>
    <li> // update a song<br/>
        PUT: <a href="#">http://localhost:5555/song/1</a><br/>
        body:<br/>
        {<br/>
	    "title": "Photograph",<br/>
  	    "singer": "Ed Sheeran"<br/>
        }</br>
    </li>
    <li> // delete a song<br/>
        GET: <a href="#">http://localhost:5555/song/3</a><br/>
    </li>
</ul>
