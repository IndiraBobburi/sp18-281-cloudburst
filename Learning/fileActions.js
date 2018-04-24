//var mysql = require('./mysql');
var glob = require('glob');
var mongo = require("./mongo");
var mongoURL = "mongodb://localhost:27017/dropbox";

function getData(req,res){
    getFiles(req.params.userid,res,'0');
}
function upload(req,res){
    
    var dateTime = require('node-datetime');
    var dt = dateTime.create();
    dt.format('m/d/Y H:M:S');
     mongo.connect(mongoURL, function(){
        var myobj = {
            "fileid" : 1,
            "filename" : req.file.originalname,
            "filetype" : 0,
            "owner" : true,
            "star" : false,
            "activityIndicator" : true
        }
        var myquery = { userid: req.params.userid };
        var newvalues = { $push: { data: myobj } };
                var coll = mongo.collection('user');
                coll.updateOne(myquery, newvalues, function(err, mongores) {
    if (err){
        	res.status(401).json({data: err});
    }
                    else{
                        console.log("1 document inserted"  + res);                        getFiles(req.params.userid,res,req.params.currentFolder);
                    }
  });
            });
}

function getFiles(userid,res,currentFolder){
     mongo.connect(mongoURL, function(){
                var coll = mongo.collection('user');
                coll.findOne({userid: userid}, function(err, user){
                    console.log(user)
                    if (user) {
                        
                       res.status(201).json({mongores: user.data});
                        

                    } else {
                       res.status(401).json({mongores: err});
                    }
                });
            });
}
function starFile(req,res){
    console.log("in star file");
    mongo.connect(mongoURL, function(){
        var myquery = { userid: req.params.userid };
         console.log("body");
         console.log(req.body);
        var newvalues = { $set: { data: req.body.data } };
                var coll = mongo.collection('user');
                coll.updateOne(myquery, newvalues, function(err, mongores) {
    if (err){
        	res.status(401).json({data: err});
    }
                    else{
                        console.log("1 document inserted"  + res);
                        getFiles(req.params.userid,res,req.params.currentFolder);
                        
                    }
    
  });
            });
}
function createFolder(req,res){
    var folderName = req.params.name;
    var inFolder = req.params.inFolder;
    var resultspre;
    mongo.connect(mongoURL, function(){
        var myobj = {
            "fileid" : 1,
            "filename" : folderName,
            "filetype" : 1,
            "owner" : true,
            "star" : false,
            "activityIndicator" : true
        }
        var myquery = { userid: req.params.userid };
        var newvalues = { $push: { data: myobj } };
                var coll = mongo.collection('user');
                coll.updateOne(myquery, newvalues, function(err, mongores) {
    if (err){
        	res.status(401).json({data: err});
    }
                    else{
                        console.log("1 document inserted"  + res);
                        getFiles(req.params.userid,res,req.params.currentFolder);
                    }
    
  });
            });
    
}
function shareAction(req,res){    
    var resSharedTo = req.params.userid;
    var resid = req.params.resid;
    var restype = req.params.restype;
    var userid;
    
    
    /*var getUser="select * from users where emailid='"+resSharedTo+"' or firstname='" + resSharedTo +"' or lastname = '" + resSharedTo +"'";
     console.log(getUser);
     mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
         console.log(results);
               userid = results[0].userid;
          var shareQuery;
    if(restype === 'file'){
        shareQuery = "INSERT INTO `dropboxdb`.`user_files` (`userid`, `fileid`, `star`, `owner`) VALUES ("+userid+ "," + resid + ",'false','false')";
       }
    else{
         shareQuery = "INSERT INTO `dropboxdb`.`user_files` (`userid`, `star`, `folderid`, `owner`) VALUES ("+userid+ ",'false',"+resid+",'false')";
    }
    console.log(shareQuery);
     mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var obj={
               "data" : results
           }
           res.status(200).json(obj);
               },shareQuery);
           //res.status(200).json(obj);
               },getUser);*/
    
    
   
}
function deleteAction(req,res){
    
   var userfileid = req.params.userfileid;
   var userid = req.params.userid;
   var type = req.params.type;
    var fileidtemp = req.params.fileid;
    var deleteQuery;
    var validDelete;
    if(userfileid !== null){
       console.log("hvhj" + userfileid)
    var validDeleteQyery;
    validDeleteQyery = "SELECT * FROM `user_files` where user_files.userfileid =" + userfileid;
    
    mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   /*var obj={
               "data" : results
           }
           res.status(200).json(obj);*/
        validDelete = results[0].owner;
        var fileid = results[0].fileid;
        var folderid = results[0].folderid;
        if(validDelete === 'true' && type === 'file'){
            
            
             var dateTime = require('node-datetime');
    var dt = dateTime.create();
    dt.format('m/d/Y H:M:S');

    var activityQuery = "INSERT INTO `dropboxdb`.`user_activity` (`userid`, `event`, `eventtime`) VALUES("+req.params.userid
    +",'Deleted the file "+ req.params.filename
    +"','"
    +new Date(dt.now())
    +"')";
    console.log(activityQuery)
                  mysql.fetchData(function(err,results){
               },activityQuery);
            
            
            
            deleteQuery = "DELETE FROM `dropboxdb`.`user_files` WHERE user_files.fileid = "+fileid;  
             mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                  
                   var deletefilesQuery = "DELETE FROM `dropboxdb`.`files` WHERE files.fileid = "+fileid;
                  mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var obj={
               "data" : results
           }
           res.status(200).json(obj);
               },deletefilesQuery);
               },deleteQuery);
        }
        else if(validDelete === 'true' && type === 'folder'){
            
            
            
            
              var dateTime = require('node-datetime');
    var dt = dateTime.create();
    dt.format('m/d/Y H:M:S');

    var activityQuery = "INSERT INTO `dropboxdb`.`user_activity` (`userid`, `event`, `eventtime`) VALUES("+req.params.userid
    +",'Deleted the folder "+ req.params.filename
    +"','"
    +new Date(dt.now())
    +"')";
    console.log(activityQuery)
                  mysql.fetchData(function(err,results){
               },activityQuery);
            
            
            
            deleteQuery = "DELETE FROM `dropboxdb`.`user_files` WHERE user_files.folderid = "+folderid; 
                mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var deletefolderQuery = "DELETE FROM `dropboxdb`.`folders` WHERE folders.folderid = "+folderid;
                  mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var obj={
               "data" : results
           }
           res.status(200).json(obj);
               },deletefolderQuery);
               },deleteQuery);
            }
        else{
            var obj={
                messge:"mot authorized to delete"
            }
            res.status(200).json(obj);
        }
        
               },validDeleteQyery);
       }
    else{
         var deletefilesQuery = "DELETE FROM `dropboxdb`.`files` WHERE files.fileid = "+fileidtemp;
                  mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var obj={
               "data" : results
           }
           res.status(200).json(obj);
               },deletefilesQuery);
    }

}
function  getuseractivity(req,res){
    var activityQuery = "SELECT * FROM user_activity where userid = " + req.params.userid;
      mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var obj={
               "data" : results
           }
           res.status(200).json(obj);
               },activityQuery);
    
}

function getFolderData(req,res){
    var resilt1;
    var folderid=req.params.folderid;
    var query1 ="SELECT * FROM files  where dropboxdb.files.folderid = " +folderid;
          mysql.fetchData(function(err,results){
                     if(err) {
			res.status(401).json({data: err});
		}  
                     resilt1 = results;
              
              var query2 ="SELECT * FROM user_files  INNER JOIN folders ON `folders`.`folderid` =`user_files`.`folderid` where `user_files`.`userid` = "+req.params.userid+" and folders.infolder = "+folderid;
              console.log(query2)
               mysql.fetchData(function(err,results){
                    if(err) {
			res.status(401).json({data: err});
		}   
                   var obj={
               "files" : resilt1,
               "folders" : results
           }
           res.status(200).json(obj);
               },query2);

                     }, query1);
}

function unstarFile(req,res){
    var userfileid = req.params.userfileid;
    var unstartFileQuery="UPDATE `dropboxdb`.`user_files` SET `star` = 'false' WHERE `userfileid` ="+ userfileid;

    mysql.fetchData(function(err,results){
        if(err){
			res.status(401).json({data: err});
		}
		else 
		{
           res.status(200).json(results); 
        }
    },unstartFileQuery);
}

exports.getData = getData;
exports.getuseractivity = getuseractivity;
exports.upload = upload;
exports. getFiles =  getFiles;
exports.starFile = starFile;
exports.unstarFile =unstarFile;
exports.createFolder =createFolder;
exports.getFolderData =getFolderData;
exports.shareAction = shareAction;
exports.deleteAction = deleteAction;
