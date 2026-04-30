
Transkription gestartet

  0:03
meiner Arbeit pseudo randomisiert, also eure Namen kommen jetzt nicht direkt vor. Wenn ihr wollt, kann ich euren Namen in meinem Acknowledgements erwähnen, wenn ihr nachdem ihr ja einen gewissen Teil an meinem Research beigetragen habt. Genau, ich werde Sachen zu Secretmanagement und zu eher privateren Prozessen einer Organisation
Stellen, die müsst ihr nicht beantworten / könnt ihr so oberflächlich beantworten, so dass halt nichts, was nicht sein sollte oder nicht veröffentlicht werden sollte, herauskommt. Aber wer sieht die *** und die ganzen Business Units werden nicht genannt in meiner Forschung, weil ich das bei Sehe als ein.

  0:37
Mhm.

  0:39
Also, ich habe die ganze Forschungsarbeit nicht auf die Firma aufgebaut, sondern im generelleren Konzept, weil das ja einfach schön generell Framebares ist. Genau, und sonst halte ich einmal ein paar organisatorische Fragen an euch alle, dass ich sozusagen dann euren Expertenstatus seiten kann.
Und dafür müsste ich von jedem von euch wissen: Wie lange macht ihr das, was ihr macht? Schon was ist eure genaue Rolle, sozusagen? Und beschäftigt ihr euch täglich/wöchentlich mit Application Security Management? Und was haltet ihr vom Approach Maximal Secret Operator?
Besten Fall jetzt jeder nacheinander. Das ist halt mit dem Dreiertermin jetzt ein bisschen overhead, aber sonst der Rest passt.
Von der Struktur für mich sehr gut.

  1:21
Ich fang mal an. Ich weiß nicht mehr alle Fragen, die du jetzt gehabt hast, aber erste Frage. Ich bin jetzt 16 Jahre in der ***, war 13 Jahre lang im *** Team tätig.

  1:23
No.

  1:37
als Softwareentwickler, Architekt, Senior Softwareentwickler, Produktverantwortlicher, alles irgendwie ein bisschen, war relativ kleines Team. Genau, *** ist, weiß nicht, kennst du wahrscheinlich, der eigenentwickelte Identity Provider.

  1:50
Yeah.

  1:56
Genau, und bin jetzt seit 3 Jahren beim *** im Team im Application Security Team und wir beschäftigen uns tagtäglich mit Sicherheit von Softwareanwendungen. Welche Vorgaben gibt es? Wie können wir die umsetzen?
Can I know.

  2:14
Wunderbar, ja gut, dann heißt es, ihr seid alle wahrscheinlich mit Secret Management beschäftigt jeden Tag. Das brauch ich dann nicht jeden Einzelfragen und ja, ihr gut, ihr seid doch nicht zufrieden mit dem aktuellen Approach, sonst würdet ihr ja nichts machen. Genau und wie gesagt, Secret Operator Kontaktpunkte hat jeder von euch schon.

  2:21
Genau.

  2:33
Oder ist das Konzept jemanden neu?

  2:34
Ich, ich kenne nur verwendet an sich, haben wir noch nicht.

  2:34
Ich noch nicht, Mann.

  2:40
Passt, dann werde ich dazu noch kommen. Gut, dann noch von dir, ***. Wie lang machst du das schon und was ist deine Rolle?

  2:41
Mhm.
Ich bin seit 5 Jahren in der *** im Platform Application Security Team. Meine Rolle ist genauso auch SSDLC Technical Standard Ausarbeitung, aber vor allem auch Security Training in der ***.
Ja.

  3:08
Wunderbar, gut. Dann noch von dir, ***, und dann können wir schon mit dem spannenden Teil beginnen.

  3:15
Genau, ich bin seit ungefähr 15 Jahren in der *** Market Application Security seit.
67 Jahren. Ich weiß es leider nicht ganz genau, wann wir gestartet haben. Wir haben begonnen, das Team aufzubauen und sind nach und nach sind ein paar Mitarbeiter da dazu gekommen und
Ja, die also wir sind eher jetzt die mir sind eher die Anforderungen an Security Management jetzt bekannt als eine Implementierung an sich. Deswegen ist gut, wenn du das mal vorstellst.

  3:56
Wunderbar, gut. Also, dann beginnen wir mal mit ein paar Konzepten, die wir jetzt wirklich brauchen werden. Ein Centralest Secret Store, das es sozusagen heutzutage Systeme gibt, die sich um den Lifecycle von Secrets kümmern, wo ich Revisions habe, wo ich die zentralen Speicher, was eigentlich eine REST API ist, die sehr geharden, die ist fürs Retrieven von Credentials, ist uns allen im Begriff.
Da glaube ich, brauche ich nicht tiefer reingehen. Genau, aber sozusagen einer der komplexesten oder interessantesten Aufgaben zur Zeit ist eigentlich dann diese Werte in eine Applikation reinzubekommen. Und je nachdem, ob du das, wenn Diagnostik oder von einem Wenn da abhängig machen willst, hast du halt verschiedene Wege.

  4:17
Mhm.

  4:32
Ein Weg, den ich als sehr interessant und effizient finde, ist der X and a Secret Operator. Dem kann ich nämlich eine Definition einer Verbindung zu einem Secret Store sowie ein Secret Template geben. Dieses Secret Template löst diese Operator dann auf mit den credentials, die er über die Store Connection hat.
Und bildet daraus dann sozusagen native Kubanetic Secrets. Das heißt, der ist eigentlich auf gut Deutsch ein Proxy, der einfach mit dem Secrets Store redet, mit der Authentication, die konfiguriert wurde und die das dann in Kubanetic Secrets rendert.
Genau, und diese Kubernetes Secrets sind dann mountable bei Pods wie alles andere, was man aus der Kubernetes Welt kennt. Was das natürlich wieder relativ cool macht, ist, dass diese Secrets eine Time-to-Live haben. Das heißt, die werden automatisch geupdatet, solange der Operator halt Connections zum Secret Store hat.
Und er unterstützt viele Secret Stores, was sozusagen dafür sorgt, dass du, wenn Diagnostik unterwegs sein kannst und sogar mehrere Stores pro Cluster Namespace angeben könntest. Das heißt, du bist jetzt nicht forciert, jetzt ewig lang ein Hashikovolt zu verwenden oder ein Azure Vault. Du könntest dich relativ einfach.
migrieren, ohne dass du irgendwas an deiner Applikation ändern musst. Und was auch noch super toll ist, dass du Files hast, hast du keine Runtime Dependency zu diesem Agent, weil sobald das File geschrieben wurde in Cubanetes, was ja dann als Secret da ist, starte die Applikation. Natürlich gibt es Applikationen, die sagen, okay, eine Runtime Dependency zu Asia ist supercool.
Das stört uns nicht, weil bei uns läuft alles in der Cloud. Das sind halt nicht alle Applikationen. Das heißt einmal, wenn ihr eine Applikation hat, die wichtiger ist, dass sie hochfährt und die Datenbankpasswörter eine Nacht olliert werden, bringt es dir nichts, wenn die Asia nicht da ist, dass deine Applikation nicht geht. Genau, und du kannst natürlich natürlich auch.
Brownfield oder alte Systeme inkludieren, die sozusagen jetzt nicht die neuesten Standards haben. Genau, aber das ist eigentlich ein super tolles Konzept, aber es gibt kein Tool, keine CLI, nichts, was sozusagen so konzeptionell gleich funktioniert auf Linux Agents.
Genau darum habe ich mir sozusagen basierend auf dem External Secret Operator einen Workflow überlegt, der sozusagen das verbindet. Das heißt, es gibt dann den External Secret Operator für die Kubernetes Workloads und einen Secret Agent, den ich gebaut habe, der sozusagen die gleiche Funktionalität hat.
auf Unmanaged Linux Hosts abbildet, der dir auch die Möglichkeit gibt, deklarativ Secret Stores und Secrets zu definieren, die er dann als Files rendert und am Host ablegt. Genau. Konzeptionell ist, ich checke meine Secrets und Secret Configuration an. DCI CTI Pipeline Application, je nachdem, ob's Co
Bethesis oder den Linux Host.
Und der und der wie es und der der Agent oder der in beiden Fällen dieser Agent Prozess geht dann zum Secret Store, holt sich die Werte, template die die und macht den Detail daraus.
Genau. Jetzt kommen wir noch ein bisschen genauer zu meinem aktuellen Solution Approach. Das heißt, wir haben dann ein kleines Binary, was auf jedem Host läuft. Dieses Binary exposed ne REST API. Er hat gebundled das SQLite Binary und baut sich sozusagen lokalen SQLite Datastore auf.
Um halt die Secret Templates sowie die Connection Secrets unentcryptiert abzulegen. Das ist einer der Schwachstellen vom Prototypen zur Zeit noch. Genau und auf jeden Fall, so wie sozusagen bei Kubanetis die Sachen halt in dieses ETCD reingelegt werden, legt der sich halt in seine lokale SQLite Database.
dann kriegt er halt diesen dieses Template, template jetzt dir und aufgrund deiner Time to live refresh das halt die ganze Zeit, dass du halt sozusagen immer auf dem Stand bist, wo du sein möchtest. Dexter Secret Operator sowie mein Tool hat das eine Problem, wenn es einen Breach gibt.
Müsstest du sozusagen die Secrets Rerollen und dann einen kompletten Roll Rollout deiner Secrets triggern, wenn deine Rerolltime relativ lang ist. Also, wenn du 12 Stunden Rerolltime einstellst und du deine Secrets Rerollst und der hattest vor kurzem Rerolled, müsstest sozusagen alle Secrets nochmal.
anstoßen. Das ist sozusagen eine Limitierung in dem Approach. Aber ja, genau, was macht er? Er schreibt dann diese Datei auf die konfigurierten Hostpath, schaut halt, ob die Datei auch wirklich schreibbar ist, macht das alles und swapt sie atomically. Nebenbei exposed einen Graphaner, also ein Promifer's Endpoint, der Scrapable ist, der dir dann halt

  8:34
Mhm.

  8:51
deine Uptime garantiert sowie auf diesen Endpoint publiziere, alle Rerolls, die er gemacht hat und alle, die funktioniert haben. Das heißt, du kannst du darüber rein, dass du sofort mitbekommst, wenn irgendwas nicht funktioniert, weil das, weil keine Ahnung jegliches.
Sagen wir 4 Environments, hast du schon mal 4 Kisten, also 4 FVMs, die halt alle sozusagen irgendwie gemanaged werden müssen, wo überall was fehlschlagen kann. Demnach ist Fail Silent halt keine Option für dieses System. Ja genau, das ist sozusagen einmal das Konzept des Agents.

  9:12
Mhm.

  9:22
Dazu mal fragen.

  9:24
Warum an Restedpoint? Wofür ist der gut?

  9:28
Weil das sozusagen konzeptionell das Gleiche ermöglicht wie Kubanetes. Bei Kubanetes ist ja die Control Plan, wo du deine Commands dagegen schickst, deine deine Jammel Files. Und das Gleiche bietet dieser Agent auch. Das heißt, die ganze Konfiguration von wann, welche Secret.

  9:37
Mhm.
Ah, OK.

  9:42
Wird über diese REST API abgebildet.

  9:45
OK.

  9:46
um halt sozusagen eine C. I. C. D. zu haben, die einfach mit entweder Cube C. D. L. oder Curl Befehlen dir das erlauben kann. Genau, Secret Stores noch mal, das hab ich eh schon angesprochen, ist sozusagen eigentlich nur 'ne Konfiguration, die dir aussagt,

  9:59
Mhm.

  10:02
wo ist mein Store, wie wie connect ich mich zu denen, mit welchen Secrets. Zurzeit ist das alles sehr rudimentär mit einfach den Connection Details reingeschrieben. Könnte man aber, wenn man das auf weiter ausbaut, auch mit Managed Identities und was weiß ich alles ausstatten.
Aber für den konzeptionellen Workflow ist es halt einfach eine YAML-Datei, wo drin steht, wo ist Avault und wie authentifiziere ich mich bei dem.
Genau, und den Managed File ist dann sozusagen eine Secret Definition. Die kann ich auch über den Rest Endpoint appliane. Da gebe ich sozusagen im Templating an. So schaut meine Datei aus und diese Variable hol bitte aus dem Secret Store, hat unsere Rotation Policy, von der ich schon geredet habe.
Und hat sozusagen ein Output-Format. Output-Formate wie JSON, YAML, N von PEM-Files werden sozusagen validiert, ob das valide Files sind. Denn rein theoretisch das hat der XL Secretarator zum Beispiel nicht. Wenn ich zum Beispiel ein JSON habe, was ich templaten möchte.
Aber irgendjemand gibt da jetzt ein Kotation Mark ein, dann hättest du eigentlich eine kaputte Datei, die generiert wird. Das heißt, dass das sozusagen macht er halt mit einem kleinen Marshalling, schaut er sich halt das an, ob das passt als Datei und er schreibt mit 0 600 oder 0 644, je nachdem.

  11:05
Mhm.

  11:05
Mhm.

  11:15
Wie man das berechtigungsmäßig am Linux-Host machen möchte. Das eine ist halt nur der User von dem Age, also mit der der Agent, der auch also von dem der Agent gestartet wird, kann es lesen. Das andere ist eine Gruppe kann es lesen, also die Gruppe, in der der Agent ist, kann es lesen.
Das muss man halt sozusagen sich so einstellen, wie man möchte. Genau so, ich muss kurz was trinken.
Genau. Nochmal kurz zusammengefasst, der Agent hat eine Krate PI für Store und Secret Management. Er hat auch einen Gate Endpoint. Das heißt, ich kann mir sozusagen ganz einfach meine definierten Ressourcen noch wiederholen. Die sind halt im Fall von den Stores retected. Das heißt, die Secret Credentials kommen niemals raus, mit denen er sich authentifiziert.
Und was auch, um sozusagen, es kommen auch die Secret Templates zurück, aber nur die Templates, nie die aufgelösten Sachen, das heißt, die verlassen niemals den Host. Genau, er hat einen Audit Log, das heißt, jede Aktion wird protokolliert, dass es halt Audible ist. Dann die Time to Live von jedem Secret, über die wir schon gesprochen haben, den Matrix Endpoint.
Und jeder Pfad, auf dem er schreibt, muss explizit laut werden, dass du halt diese CVCV 35 glaube ich, ist das Paftraversals abgefangen hast. Das heißt, das kannst du zwar auch über den System User unter Linux lösen, aber wenn dein System User jetzt vielleicht die Applikation rennen darf und den Agent,
betreibt, dann möchtest du vielleicht trotzdem, dass der nur in den Slash Confolder reinschreiben darf und nirgendwo sonst. Das heißt, das musst du da noch mal explizit konfigurieren. Du kannst natürlich auch einfach ein Slash angeben. Das ist halt einfach nur eine Security Mechanismus mehr, um sozusagen.

  12:36
Hm.
Mhm.

  12:50
da wegzukommen. Jetzt mal für den Worst Case. Gegen was kann der Agent nicht, wenn der Host komplementiert wurde und Root Level Access geholt wurde? Dann sind alle Secret Daten sowie die Connections zum Store offen. Aber das Problem hast du generell bei ganz vielen Verfahren.

  13:05
Mhm.

  13:05
Sobald Root Schrägstrich die Festplatte gestohlen wurde und die halt nicht richtig verschlüsselt wurde, sind deine Daten halt.
Nicht wirklich recoverable.

  13:15
Ja.

  13:16
Genau.
Genau, und jetzt sozusagen erkläre ich noch 2 kleinere auf Konzepte meines Agents. Der Agent sozusagen ist ja eigentlich dafür ausgelegt, dass er eine REST API nach außen hat. Demnach kann man ihn so konfigurieren, dass ein MTLS-Check gemacht wird und dann erst der JVT validiert wird.

  13:29
Mhm.

  13:34
Aus dem einfachen Grund, JVTs sind replayable in der Zeit, die wir haben. Das heißt, kann ich zwar auf 30 Sekunden einstellen, aber auch 30 Sekunden könnte vielleicht für einen für einen Tag reichen und sagen, der Key Cload gibt dir ja einen JVT.
Und das Schöne ist, den kannst du ja validieren, ohne dass du sozusagen Credentials brauchst, weil die Public, weil der Public Key zum Validieren, der ist auf den Punkt, weil no nexposed, den kann sich der Agent einfach holen, wenn du es konfigurierst und kann dann die JVTs validieren. Aber ohne das M. T. L. S. hast du halt.
die Möglichkeit, dass ein Client, der das der halt nicht das Client Zertifikat hat, sich auch authentifizieren könnte. Darum habe ich mich hier für einen Layered Approach entschieden, um ihn rein theoretisch Internet Facing betreiben zu können, mit mehr Security, also sozusagen mit einer Layered Security.
Genau. Was du aber brauchst, sind die MTLS-Dateien am Host, das heißt, die zu Bootstrappen rüber, könnte man probieren, über den Agent selber zu lösen, indem man einfach die Pfade sozusagen vom Agent selber befüllen lässt. Aber zurzeit in der aktuellen Version
Macht er das sozusagen im Ensible Kopierschritt Schrägstrich im Ensible Vault, je nachdem, wie man das halt lösen möchte. Genau, dann noch zum Atomic Filesop, genau, damit diese Filesops halt, also eine Applikation, die pendet halt erst recht, Spring Application, die penden halt auf dieser File, auf dem Fall, dass die valide sind.
wenn das Teil nicht valid ist, fährt die ganze Applikation nicht hoch und das darf eigentlich nicht passieren aufgrund eines Templating-Fehlers. Demnach schaut der als erstes, darf ich in den Pfad überhaupt rein konfigurationsmäßig, ist den Content, ist der Content, den ich bekommen habe, valide in Form von kann ich den marshen und anmashen?
Wenn ja, geh weiter, dann kreieren wir halt den Temp File. Beim Temp File setzen wir die richtigen Permissions, dann schließen wir das Temp File und stellen sicher, dass es da ist und dann atomicly swappen wir das mit einem OS. Punkt rename command, das einzige, was auf Linux wirklich atomicly files renained.
und können sozusagen damit sicherstellen, dass wenn eine Datei da ist, dass sie nicht halbert geschrieben irgendwo existiert, sondern wenn wir es, wenn sie da ist, wird sie richtig geswoppt. Genau. Wie installiert man den Agen at the Moment, nachdem man sozusagen ja eigentlich, also Quinetes ist ja unser großer Ideengeber und der Connected Operated immer
braucht mal irgendeinen Weg, der ähnlich ist. Man könnte natürlich jetzt hergehen und probieren, so was wie Talos Linux zu bauen, eine irgendetwas, wo man sozusagen alles über RES konfiguriert, aber das gibt's halt in den meisten Systemen nicht. Demnach hab ich mich für Ensible entschieden, weil's finde ich die beste Alternative dafür ist. Das demnach gibt's ein Playbook.
Das kreiert ein System User, eine Secret Gruppe, das das kreiert die Pfade, die er braucht. Dann templated es dir die Konfiguration des Agents, macht das System D-Service und lädt das Binary vom GitHub runter. Genau, also ein relativ einfacher Install Installweg.
Uninstall ist natürlich das Reversen von dem, ist auch rein theoretisch dabei. Jetzt nichts allzu Komplexes. Die Konfiguration des Agents, ich kann einen Namen geben, dann die Trusted Roots, von denen ich gerade schon gesprochen habe, dann halt einen Port, für den Moderest Server dann erreichbar sein wird.
Dann kann ich die Logging Sachen konfigurieren. Ganz wichtig, dass ich den Audit Log Pfad angebe, dann wo er das Database File hinschreiben soll. Dann kommen wir auch schon zum Spanner darin, zum Security Part. Hier kann ich halt TLS konfigurieren. Und genau da kann ich halt auch das Client.
File angeben und damit kann ich sozusagen das normale TLS sowie MTLS konfigurieren und dann kommen wir zur Authentication. Da kann ich mich halt zwischen Non bis JVT für alles entscheiden. Am besten ist, wenn man halt das Client CE Files setzt, weil dann macht dann MTLS Check.
Und verwendet den Pass JVT dann für die Authorization. Man könnte auch nur MTLS machen, dann würde halt die MTLS Credentials für die Authentication nehmen. Da muss man halt wissen, was man in welchem Raum an Sicherheit braucht. Genau, ich habe mich auch ein bisschen mit SP von 2 Jahren rumgespielt. Keine Ahnung, ob euch das was sagt.
Das sind Workload, das ist ein Open Framework für für Workload Identities, wo am Host dann eine A.P.I. vorhanden ist, um Short Lift X.509 Zertifikate zu generieren. Red Hat macht da ihre eigenen Sachen, aber das ist out of den Scope von dem aktuellen Research, das heißt.

  17:24
Mhm.
Mhm, OK.

  17:42
Ich hab mich, man könnte es konfigurieren, testweise, aber bitte nicht unendlich viel bringen. Ich wollte es nur erwähnen, dass man sozusagen in die Richtung auch weitergehen könnte. Genau, dann würde ich sagen, kommen wir noch zu 2 Configuration Examples.
Und zwar einen Asia Vault. Hier sehen wir es, wie es im externen Secret Operator ausschauen würde. Hier konfigurieren wir die Tenal ID, die Vault URL und dann in dem Fall Client ID und Client Secret. Das Gleiche würden wir halt sozusagen bei meinem Agent konfigurieren, dass wir den Typen sagen.
die URL, die Client ID, Client Secret and Tenant ID, genau. Und die Daten liegen dann halt unecrypted in der SQLite Database am Host herum. Also ist halt eine Security Limitation, die es zur Zeit gibt, die man sozusagen bedenken sollte.
Jetzt kommen wir auch zum spannenderen Teil. Wie schaut denn so ein Secret aus? Auf der rechten Seite sehen wir so ein external Secret Operator. Was wichtig ist, was ich noch vergessen habe: Es gibt halt hier einen Namen für den Secret Store.
Hier ist es der Name. Hier ist es der Name.
Genau und sozusagen, dann kann ich hier einfach sagen, mit den Connections, Details von dem von dem Security genannt, bitte lös mir dieses Secret auf und hier template ich halt in was auch immer ich möchte und wie auch immer ich das möchte herum.
In dem Fall jetzt einfach ganz normales User Passwort, die URL von der Database von einem Kontext XML. Genau, ihr beide verwenden den normalen Goling Templating Standard. Das heißt, ich habe hier Punkt und hier habe ich ein Mapping von hier im Template.
Zum Secret Store das Gleiche habe ich hier drüben beim External Secret Operator auch da habe ich ein Mapping von hier auf wie heißt zum Secret Store genau und je nach Vendor schauen die Pfade halt ein bisschen anders aus, weil zum Beispiel im Haschikop wollt, kann ich mit Slashes Hierarchien im Pfad Namen angeben.

  19:17
Mhm.

  19:30
Das geht beim Asia Keyboard nicht, weil der ist halt dafür da, dass du mehrere Instanzen hast. Während Haschikop wollte für da ist, du betreibst eine Instanz und hast mehrere Paws mit mehreren ACLs. Das ist halt ganz organisationabhängig. Es gibt noch unendlich unendlich andere.
Mein Agent zur Zeit unterstützen Azure Keyboard und den Hashikob Vault. Das Hinzufügen ist nicht wirklich komplex. Das ist ein Interface and Go, was man satisfyen muss. Genau.
Ja, und jetzt hätte ich einen Server für euch, was ich euch bitten würde zum Ausfüllen und dann oder davor noch gerne Fragen, Anmerkungen, irgendetwas zu dem ganzen Thema. Ihr wart sehr still währenddessen.
Und sonst habe ich dann ein paar offene Fragen am Ende, ja.

  20:08
um
Um.
Die Fragen soll ich wieder vergessen.

  20:18
Hättest mir gern zwischendrin rein.

  20:20
Genau.

  20:21
Bevor du beginnst, ***, ganz grundlegende Frage, Maximilian, vielleicht hast du vorher schon gesagt, was war dein, was war dein Treiber genau für die für diese Implementierung, war das jetzt, ist das aus der bayerischen Informatik herausgekommen oder aus deinem Studium und du hast es nicht genutzt, weil es zufällig gepasst hat?

  20:25
U.

  20:39
Wie, wie ist denn da das zustande gekommen? Und kurze Erklärung dann noch.

  20:45
Mhm, also.
Das Ganze hat angefangen, als ich mit dem Master begonnen habe. Habe ich angefangen, DevOps zu sein bei uns und da kam das Secret Thema auf. Und dann haben wir uns ein bisschen im Kreis gedreht, was wir jetzt eigentlich genau haben wollten. Und das, was sozusagen wir nicht wollten, war aus architektonischer Sicht in einer Runtime Dependency zu Asia.

  20:54
Mhm.

  21:05
Order zum Secret Store.

  21:08
Mhm.

  21:09
Ja.

  21:09
Dann bleibt dir relativ wenig übrig

  21:12
Mhm.

  21:13
Entweder du baust dir selber dein Glue Code Skript, was das für dich erledigt, oder du verwendest halt Excel als Secret Operator und nachdem.

  21:20
Okay, also es kommt konkret aus der *** jetzt raus, sozusagen, und dann

  21:24
Es ist das, was die OpenShift-Admins für Cubanedis empfehlen, weil sie das selber betreiben und verwenden.

  21:31
No.

  21:32
Mhm, doct spezielles af Down Kedafe. Mhm, ja, passt.

  21:32
Us the idea.
Ja genau, also per se, du kannst, dass es schön ist, dass bei dem Ding dazu kommen, bei den offenen Fragen, dann auch. Du kannst eindeutig jegliche Application betreiben, solange sie es von Files akzeptiert. Und solange du dieses File watched, kannst du sozusagen entweder mit den Daten dann
welche neu konfigurieren at Runtime oder deine Applikation neu starten lassen. Das musst du halt dann wissen, wie du das machen möchtest. Der Agent gibt per se keine Signals oder irgendwas an die Applikationen, einfach aus dem Grund, dass der Agent keine Abhängigkeit zu den Applikationen haben möchte. Der möchte nichts von denen wissen, der möchte einfach nur Agents machen, der möchte einfach nur
Files generieren und wenn die Files generiert hat, dann kann die Applikation damit das machen, was sie möchte. Genau.

  22:19
Mhm.

  22:21
Antworte: Ich glaube, du solltest die Frage gut beantworten.

  22:24
Danke.

  22:26
Um.
Wie ist denn geplant, das Client Potential oder halt das SIG das?
Das Secret zum Zugriff auf den Keyboard in das in den Jason File zu bekommen, auf der auf der in in der in der Kybernetis Word referenziere ich dort ein Secret, um das kümmere ich mich irgendwie.

  22:44
Die Sache ist.
Yeah.

  22:52
Muss ich jetzt da zum Beispiel meine Deployment Pipeline, muss ich halt irgendein Templating haben, weil ich kann den File nicht einchecken mit Sigrid. Das heißt, die müsste in meiner Pipeline ein Templating haben, damit ich den File on the fly.

  22:57
Genau.
Genau.
Genau.
Ja, oder ich mache es einfach so auf dem gleichen Weg, wie ich meine Kubanetic Secrets generiere, generiere ich auch das. Weil, wenn man ganz ehrlich ist, diese sozusagen diese Referenzen hier sind cool und so, aber wie viel bringen sie dir wirklich? Also, wie viel bringt dir dieses File im normalen Git Repository?
Solang der A also mein Agent hat halt ein Get Endpoint, da kriegst du alle Stores, die es pro Haus gibt.

  23:26
Genau.

  23:30
Das heißt per se, ob du dieses File wirklich einchecken musst, um glücklich zu sein, das ist halt eine Use-Case-Frage.

  23:37
Genau, aber ich hätte im OpenShift oder im Cubanitis zum die Möglichkeit, ich leg diesen File einmal manuell an und dann ist er da. Ist jetzt natürlich Infrastructures Code nicht zwingend, aber dann habe ich nur ein Sigrid, wenn alles abgehackt, dann muss ich halt das einmal manuell neu erzeugen, da eintragen.

  23:41
Ja.
Genau.

  23:56
Dann wäre es da die Möglichkeit, habe ich auf der anderen Seite einfach gar nicht, dass ich am Linux-Server irgendwas konfiguriere. Und wie ist die Trennung?

  23:59
Genau, noch nicht.

  24:06
Sag mal, wir haben auf einem Linux-Server unterschiedliche Anwendungen von unterschiedlichen.
Systeme laufen, also von unterschiedlichen Teams laufen. Die sollen sich ja nicht, also hat das Ding Authorization oder nicht. Das heißt, wenn da was von unterschiedlichen Teams lauft, würden die alles überschreiben können von wem anderen und sehen.

  24:21
Night.
Nein, noch nicht.

  24:29
Und so weiter. OK.

  24:30
Ja, du könntest rein theoretisch einfach mehrere Agents pro Host haben. Du könntest einen Agent pro Tenant haben.

  24:36
OK, das wäre eine Möglichkeit. Ja, mhm.

  24:39
Bräuchtest halt einfach mehrere System D Konfigurationen, die den Agen halt und die verschiedenen User also wenn du sozusagen so multi tenant machen möchtest, dass es wirklich sicher ist, brauchst du auch mehrere System User dafür.

  24:49
Mhm.

  24:50
Weil sonst gibt's halt immer einen System-User, der mehr sehen darf, als er sieht.

  24:53
You know.

  24:55
Aber ja, es wäre so abbildbar, rein theoretisch.

  24:55
OK, dok you.
Mhm, perfekt.

  25:06
Deine Fragen zur Dokumentation basieren jetzt vermutlich auf deiner Präsentation, die du uns gezeigt hast, oder hast du noch irgendwo?

  25:15
Ja, ich per se das Ganze ist ein offenes Git Repository.
Das heißt, es gibt per se eine komplett generierte, aber ja, es sollte eigentlich das reichen, was ich euch gezeigt habe.
Ich hab nicht gedacht, dass jemand schon dasselbe macht. Aber ja, gerne, wenn ihr irgendwelche Fragen zum Server habt, darum mach ich das in einem Termin. Stellt sie mir einfach.

  25:36
Mhm.

  25:42
Schick das Repository, sobald er es geladen hat.
Wann?
Genau, und da gibt es einen GitHub. Da gibt es einen Link, also einen Link von Autogenerator Doku über die Markdon Dateien im Docs Folder.

  26:04
Der Mega Maxl, sehr gut.

  26:06
Mega MXL, Mega Max Listen, ìberall Name, ìberall sonst ja.

  26:08
Um.
Lauft das jetzt schon irgendwas? Hast du das bei bei Carving schon irgendwo in in Einsatz?

  26:14
Ja, ja, ja, zur Zeit läuft es auf einen Tom Cat.
Und zwar auf dem hier, wenn wir hier unter Matric schauen, der zur Zeit eine Datei, die er managed, die er 8 mal rerollt, noch ganz ohne CICD Pipeline, einfach mal händisch. Wir sind gerade erst in der Testphase, ob das das Ding, ob das Ding macht, was es machen soll, aber sie macht das schon was, ja.
Seit jetzt vor Ostern haben wir ihn aufgesetzt, kurz als Testweise Implementierung.
Genau, dann können wir uns das hier auch kurz anschauen. Das liegt hier alles hier grad auf meinem User herum. Das heißt, hier hat man dann halt sozusagen seine Yummle-Datei, 'n Pfad, dann haben wir hier die Reference, dann haben wir hier sozusagen das Konfiguration. Also in dem Fall, in unserem Fall, ist halt die ganze Config, wo Secretwerte reingetemplated werden.
Aber das kann halt sozusagen auch eine extra Datei sein. Das heißt, du kannst halt jegliche Form von Datei einfach abbilden. Hast halt hier eine Referenz.

  27:09
Mhm.

  27:11
Die du dort oben dann halt auflöst und das macht der Agent für dich.
Kino.

  27:26
Mhm.

  27:27
Mhm.

  27:29
Genau, wenn wir es bei uns sozusagen dann wirklich aufsetzen wollen, dann ist halt nicht Promifer, also ein Grafana, der Stack, sondern *** und was auch immer, aber das sozusagen ist eine Referenzimplementierung für meine Masterarbeit.

  27:42
Mhm.

  27:42
Mhm, ja, wobei Beta da kann ich ja pro Meter Metriken auslesen, von dem her sollte das jetzt nicht das Thema sein, oder?

  27:49
Ja, der *** hat gemeint, dass wir das Promi-Feuers irgendwie auf irgendeinem Radar schon wieder als deprecated angesehen wurde. Also, ich habe keine Ahnung. Auf jeden Fall, er hat per se einen Monitoring Endpoint, bei dem muss man sich halt zur Zeit auch noch aufhändicaten, weil es ein, also wie gesagt,
Wenn das Ding wirklich produktionsweit oder mehr genutzt wird, dann werd ich da noch mal mehr Zeit reinstecken. Das ist natürlich ein Masterarbeitsprojekt, das heißt, das ist halt alles scoped.

  28:14
Mhm.

  28:16
Genau.

  28:19
I call.
Ich finde es, also jetzt von dem, was ich bisher gehört habe, sehr positiv, was du dann Zeit und Gedanken reingesteckt hast.
Cool, dass wir das in der Waschinformatik so auf dem Level machen. Das ist.
Gut zu sehen, dass sich jemand über das hinaus Gedanken macht, das, was man sonst vorgegeben kriegt. Das ist schon sehr hilfreich, sag ich mal.

  28:48
Ah.

  28:49
Mhm.

  28:50
Ich habe den Master jetzt Thema gebraucht.

  28:53
Mhm.

  28:53
Ich liebe Softwareentwicklung, das See und das das vereint beides relativ cool, ja.

  28:59
Rap.

  29:00
Jetzt genau, Survey, habt ihr alle schon beantwortet oder soll ich noch warten auf jemanden?

  29:05
Ach so, na, ich hab jetzt heut angefangen. Ich hab nur deine deine deine Fragen nur kurz da durchgelesen.

  29:06
Mhm.

  29:06
Uh.

  29:07
Ich hab noch net einmal angefangen.

  29:07
Ach so.

  29:07
Ich bin noch nicht fertig mit Beschreibungen.

  29:10
Du, alles gut. Ich möchte es halt wissen, nicht dass wir halt uns anschauen im Termin. Wir haben alle Besseres zu tun.

  29:16
Also.

  29:17
Na, ich glaube, du brauchst uns, wie gesagt, nicht zuschauen beim Beantworten. Oder hast du noch irgendeinen Punkt, weil ich würde das sonst einfach noch im Nachgang beantworten?

  29:25
Es es geht per se Open Crash. Ich bin mir nicht sicher, ob alle Fragen gut verstanden sind. Das heißt, ich glaube, am schlauesten wäre es, wenn wir es jetzt machen, dann stelle ich euch jetzt die offenen Fragen und wir machen die davor und ihr beantwortet die dann sozusagen offline und danach, wie ihr wollt.

  29:35
Ach so, du hast noch Fragen.

  29:36
Genau.

  29:41
Ich kenne das nur von mir. Wenn ich einen Server irgendwo bekomme, wo ich nicht gezwungen bin, das zu machen, dann mache ich es nicht.

  29:50
Ah ja, das also ich glaub die Gefahr, dass wir es jetzt nicht machen, ist jetzt nicht so, also bei mir jetzt nicht.
Ich glaub mal, das kannst du nicht durchklicken. Ist das eine Seite oder sonst ist mehr?

  30:03
Na, es ist die also so lang du halt runterscrollen kannst.

  30:06
Ach so, ja gut, kann ich nicht durchklicksen.

  30:10
Wenn Sie einfach alles auf Strongly Agree geben.

  30:13
Ja, das hilft. Das hilft dir sicher am allermeisten. Genau, und dann hast du es zum Diskutieren.

  30:21
Ja.

  30:22
Ohne Begründung, 3 Leuten hat's einfach gar nicht gefallen.

  30:26
Genau.

  30:31
Dann bin ich schon gespannt
Im besten Fall, wenn ihr zu so einer Frage dann schon irgendeine Anmerkung habt oder so, dann stellt sie einfach, weil dann sparen wir uns das dann bei den offenen Fragen.

  30:44
Hast du den Fragebogen selbst generiert?

  30:49
Also, ich habe ihn selbst geschrieben, nicht einmal generiert, aber ich habe sicher ein, zwei Fehler.

  30:52
Ja, ein Feedback zum Fragebogen: Sport. Das nächste Mal ist normal in der Mitte.

  30:56
Ja.
OK, die Sache, ich verstehe, aber 'Like Head Scale' ist halt definiert, dass du das normal brauchst. So, OK.

  31:05
Na na, dann ist das so, das ja, eh, das kann sich durchaus jemand mal festlegen, ob das in Ordnung oder nicht in Ordnung ist. Das macht das Mama viel zu selten. Das ist also da hinten, ja, weiß jetzt nicht.

  31:17
Mhm.

  31:22
Yes.

  31:22
Stehe, nehme ich mir mit. Werde ich mit meinem Masterarbeitsbetreuer diskutieren, was er wissenschaftlich davon hält.

  31:31
Na ja.
Sagst ihm schöne Grüße von mir. Die Welt sollte sich mal entscheiden. Nee, da ja eh Antwort geben.

  31:40
Du bist schon lang genug in der Beuscham.

  32:14
Kann ich bei der Authentifizierung einschränken, welcher welchen User ich hinlass?

  32:22
Nein, zur Zeit, zur Zeit ist es nicht userbased, das ist also der Gedanke war, du machst dir in deinem I.D.P. einen Rail Pro Agent, den du hast. Also es ist noch nicht, dass äh auch noch validiert, welcher User das genau ist, sondern nur ob der J.B.T. valide ist.

  32:37
OK, also tschüss.
Scheint mir recht früh.

  32:44
And and a random pro eight and months it's for and for key clock or challenge other when intra custes it machen.

  32:46
comfig

  32:50
Genau.

  32:51
Von *** zum Beispiel, also Key Lock wird wahrscheinlich nicht das System auf ewige Zeiten sein.

  32:59
A so talk in in intra consum auch mit unterschiedlichen app registrations, oder?

  33:02
Aber dann müsst ihr.

  33:05
Yeah.

  33:06
Where is this like I'll cut Simone.

  33:06
Genau.

  33:07
It sollte das Gleiche sagen.

  33:09
Mhm, stellen wir vom Management auf und.
Aufwendig vor.

  33:13
Ja.

  33:15
Wenn man einen Café anschaut, das ist ja nicht gerade wenig. Die man im Café ist egal. Na, wobei ist auch nicht egal. Ne Café hat ja also *** hat ja viele.

  33:22
Doch eigentlich.
Schon, aber die laufen alle im gleichen Tom Kit mit dem gleichen User zur Zeit.

  33:27
Produkte

  33:32
Also, alles, was an Tomcat ist, ist im gleichen User. Also, für Calvin ist das relativ irrelevant.

  33:38
Ja, aber ich bin.
manche Dinge laufen aufeinander und Homecat und auch das könnt ihr jetzt nicht mehr unterscheiden, weil der User kommt ja hin, wenn er also ich sage mal, keine Ahnung, Kaffee in Core läuft auf diesen 10 tomcats und *** läuft auf 10 anderen Homecats, dann sind das 2 unterschiedliche User, mein Dogstore ist jetzt

  33:50
Huh.
Ja, aber Dock Store.
Sind das 2 verschiedene Linux Hosts?

  33:59
Ja, Open Shift.

  34:02
Dockstore läuft nicht auf dem.

  34:02
Ja klar, aber ich könnte ja nicht, ich kann trotzdem den J. B. T. an den anderen Linux verschieben und dann nimmt er an. Ich kann ja nicht unterscheiden, wenn ich nicht im N. 3 ID unterschiedliche App Registrations hätte.

  34:10
Windows
Ja, genau.
Ja, das stimmt.

  34:17
Und das genau, also ich denke, das wäre schon ein Feature, das man braucht, in irgendeiner Form.

  34:21
Wobei du auch sagen kannst, du könntest es bei M.T.L.S. halt abbilden, indem wir sozusagen JBT kannst du einmal haben, aber die M.T.L.S. müssen halt zum Host passen. Sonst hast du halt das gleiche M.T.L.S., also das gleiche Client Certificate File auf jedem Host.

  34:37
Na ja, aber wir haben dann haben wir wieder das Problem, wie verwaltet denn das Client Zertifikat. Ich gehe davon aus, ich mach das aus einer Pipeline raus in GitLab C. I. C. D. Jeder hat das gleiche GitLab C. I. C. D. Das heißt, ich müsste in mein in meine Pipeline in mein

  34:46
Ja.
No.

  34:53
Builders Image also mein Image, das ich in meiner Pipeline laufen habe, dann plötzlich a credential reinkriegen a private a private Key, das heißt dieser credential Management woanders hin verlagert also das würde mir weniger weniger gefallen also.

  35:00
Yes.
Ja, natürlich, es löst.
Es löst nicht das Secret Zero Problem, das ist auch noch nicht das Konzept dieses dieses Researchers. Das Secret Zero Problem ist bekannt, aber das Gleiche hast du auch mit Maxxell Secorporator.

  35:08
Genau, also.
Genau, ich hoffe, das haben wir heuer.

  35:16
It's I verstanden. It's I verstanden, dass das hauptsächliche Problem, was du lösen wolltest, is this, die Abhängigkeit.

  35:24
Genau die Abhängigkeit und dass der und dass das Delivery halt nicht mehr irgendein Shell Skript ist, was bei Cronjob aufgerufen wird und vor sich hinfällt.

  35:25
Genau.

  35:25
Yeah.
Genau das.

  35:33
Mhm.

  35:33
Yeah, yeah.
Wenn man es sicher machen möchte, dann musst du den den Agent sparen, weil dann brauchst du es eh nicht, dann musst du wieder in den sauren Apfel beißen, dass nicht da sein kann, der das Service also.

  35:47
Na ja, aber auch dann, wenn ich Root auf der Maschine bin, dann mache ich ja dann Memory Dumper oder whatever. Also grundsätzlich, wenn ich Root auf der Maschine bin, haben wir verloren. Dann hole ich mir halt den Memory und schaue, was da drin steht an Credential.

  35:47
Ja.

  35:59
Ja, aber das sind zwei unterschiedliche Themen. Das vom Maximilian ist jetzt wirklich das Ziel.

  36:04
Nein, aber.

  36:07
Dass man die Abhängigkeit zu einem Cloud-Service reduziert.

  36:07
Oh, oh, ey.
Genau, na wie du sagst, da muss ich in den sauren Apfel beißen und die Abhängigkeit haben. Das ist aber laut in der Applikation, aber das hilft mir nichts, wenn jemand auf dem.

  36:13
Bisschen legitim.
Na, dass dass man auf einen Internetservice zugreift, was nicht da sein kann.
Nee, nicht dem, dem die Sicherheit an sich.
Sehe ich jetzt keine wahnsinnige Veränderung. Also, es ist es an Secrets so oder so und es gibt irgendwas Service, von dem ich konsumiere. Da ist noch ein Teil dazwischen, was die Secrets auch kennt. Also, per se wird es unsicherer. Das Handling an sich ist so gut wie es geht, aus meiner Sicht.

  36:32
Genau.

  36:46
für diesen als Trader für diese Abhängigkeit zu einem externen Service, wo sie diese Abhängigkeit verhindern möchte. Wenn ich diese Abhängigkeit, wenn man das nicht so holprig,
Priorität ist, ob das Ding da ist oder nicht, oder dann lies ich es halt das nächste Mal. Dann ist ganz klar dieser dieser Agent nicht hilfreich, weil per se das Secret Management aus meiner Sicht jetzt nicht für das 0815 für die 0815 hab ich jetzt auch nicht verbessert.
Ne, also ist jetzt nicht die Kritik an dem an dem Service, die ist absolut wichtig für die anderen Application, wo diese Abhängigkeit eben ein Riesenthema ist. So sehe ich das jetzt.

  37:32
Ja, ja, ja, auch die Sache ist, ähm, es macht es, äh, das Problem ist halt auch, wenn du jetzt sagst, Applikation hol die Secrets ab, irgendwie muss ich die Applikation authentifizieren können.

  37:32
Hör doch.

  37:46
Wie kriegst du die Credentials Application? Das kannst du halt nur bei Managed Identities machen. Das bedeutet, du brauchst irgendeine Authentication Pro, ähm, irgendeine Authentication API auf Host Level. Sei es managed, äh, sei es managed Identities, sei es eine Asia VM, die dir sagt, ich bin eine Asia VM.

  37:47
Genau.

  38:04
Über deren API, aber solange sozusagen.
Du, du sozusagen nicht irgend so eine API hast, musst du halt irgendwie händisch immer die Potentials rüberbringen, weil du sonst sozusagen, weil diese initial trust Phase halt sozusagen noch nicht wirklich, soweit ich weiß, gut gelöst ist für halt.

  38:10
Mhm.
Genau, also ich hoffe, freier lösen Sie es mit unserer OpenShift OnPrem. OpenShift unterstützt ja Federate Identity und dann kannst du ohne also mit SIO credential SIGates kannst du dann.

  38:33
Genau.

  38:39
Auf ein Azure Keyword zugreifen und die einfach.

  38:42
Genau das, das könnte also sozusagen, wenn das einmal da ist, das ist halt doch dieses Spy von Spire, das sind sozusagen die Open-Source-Gegenspieler zu diesem Managed Identity Dings. Das Problem ist, da hast du auch einmal einen Root Token, der rüberkommt, der rüberkommt.
Es ist sicherer als das aktuelle, aber du bist nie wirklich Secret Less, nicht einmal mit den Ferret Identities. Es gibt eine Bootstrap-Phase, wo es einen Bootstrap-Key gibt.
Von dem Research, den ich betrieben habe.

  39:11
Also, ich weiß jetzt nicht, wie es funktioniert, aber Asia ist ja so, das Arbeit mit JSON Webtokens und das Einzige, was du sicher verwalten musst, ist du stellst den JSON Webtoken aus und signierst den, den Private Key musst du sicher verwalten. Aber dann also, der Asia und mein System tauschen keine Tokens aus.

  39:25
Genau.
Ja, ja, nein, ich rede doch nicht, ich rede doch nicht von Web. Ich rede generell vom Host, aber das ist ein technisches Detail. Das können wir gern wann anders weiter diskutieren.

  39:31
Genau.
Genau, also man arbeit für.
Für mich als User wird es besser. Ich muss mich um keine Cadentials kümmern. Irgendjemand muss den Server sicher betreiben, aber das muss er generell. Genau, also das ist nicht mein Thema, genau.

  39:46
Solar, ja genau, solang du sozusagen den, solang, solang du halt den Server managed kaufst, dann schon. Aber das Tool richtet sich halt an Linux-Systeme, die nicht managed sind, wo du sozusagen keine Authentication A. P. I. hast.
Weil, sobald du halt dir eine VM bei der Asia nimmst, hast du ja die Asia Authentication API, über die du sozusagen deine Identity als Host hergeben kannst.

  40:12
Also ihr redet aus unserer Entwicklungssicht von, wir bekommen von, also wir haben P.X.I., unsere Infrastrukturabteilung, ich bekomme von P.X.I. einen Server und ich kümmer mich nicht darum, dass das da konfiguriert ist. Ich bekomme ein OpenShift gemanagt und dort ist Federal Contentials eingestellt, ich muss in meine

  40:22
Ach so, ja genau.

  40:31
Gamespace nichts konfigurieren, dass der mit Asia quatschen kann. Das ist aus Entwicklungssicht, aus Infrastruktursicht ist natürlich was anderes. Und wenn jetzt jemand seinen eigenen Linux-Server betreiben möchte in der Entwicklung, soll er nicht.

  40:33
Yeah.
Genau.
Ja, ja.

  40:47
Aber wenn, dann soll er sich darum kümmern, ja.

  40:53
Genau, gut, dann glaub ich.

  40:54
Kids to.

  40:55
Was meinst du mit der Frage Developers can manage their own secret configuration without being exposed to secret values?

  41:01
Nachdem du sozusagen nur immer das siehst, als die Welt, also insofern ein Developer keinen Zugriff auf den Server hat, wird er immer nur das sehen und diese Values, bei Si, die im Store dahinter liegen, wird der Developer niemals sehen. Das heißt, der könnte sozusagen nur auf dem Template.
Sachen ändern, aber wird sozusagen niemals die Werte kennen. Das heißt, wenn du ein Long Lift Potentials hast, so jetzt zum Beispiel von einem Ford Party System, wo man sich keine Ahnung, per Unique oder so anmelden könnte, dann müsst was der Developer jetzt könnte, halt im aktuellen Approach, wer halt.
Das, wenn am Server die Rechte die Dateien auslesen.
und mit dem so wie es zur Zeit ist und wenn du das halt abdrehen würdest, dann könnte er hier Sachen ändern und wäre niemals zum Secret Value Exposed.

  41:49
Mhm.

  41:50
Darum geht es.
Ich hätt gern eigentlich immer noch 15 Minuten insgesamt für die offenen Fragen. Wie schaut's bei euch aus?

  42:38
Strongman.

  42:38
Sie mittendrin.
Passt.

  42:41
Schmier.

  42:41
Morgen, nun gleich.
Also, wir können die offenen Fragen machen und dann machen wir das noch fertig.

  42:49
Besteht.
Was für gibt es irgendeinen Vorteil, den irgendjemand von euch sieht an diesem Secret Agent?
Über das, was du jetzt natürlich schon gesagt hast, *** hinaus, das werde ich natürlich verwenden.

  43:05
Mhm.

  43:06
Also das Protokoll, die Abhängigkeit zu verhindern oder verringern.
So weiter.

  43:13
Wunderbar.

  43:13
Fortan.

  43:15
Also die Abhängigkeit zu verhindern und durch die.
Einfachheit der Verwendung, sag ich jetzt mal, verstehst du, dass es relativ einfach ist zu verwenden für mich als als es ist DevOps Team, die Secret Management auch richtig zu machen und ich die Secrets einzuchecken.
sondern wirklich eine gute und sinnvolle an an einem Volt zu verwenden, der dafür gemacht ist und den dann anzubinden, weil es jetzt nicht mehr viel komplizierter ist als.
Die Secrets einzuchecken und dann ein Config file zu haben mit allen Secrets, auch in ein extra Report oder whatever. Ja, das sehe ich einfach jetzt als großen Vorteil.

  43:51
Ja.

  43:59
Mhm.

  44:00
Steh.

  44:01
Ja, in dem Zusammenhang, das was du jetzt vielleicht auch gemeint hast, das Thema externe library, die dir ein Problem löst und wo sich jemand anders drüber Gedanken gemacht hat, also der Entwickler oder die Entwicklerin selbst.

  44:01
Genau.

  44:17
Braucht sich über das high sophisticated Thema nicht mehr Gedanken machen, sondern nur mal diese Library zu verwenden. Das ist durchaus ein Vorteil.

  44:27
Ja, dass das ein standardisierter Weg ist und nicht jede Applikation sich selber drum kümmert.

  44:28
But you misault.

  44:28
Hm.
Mhm, ja, wunderbar. Sag ich danke für den Input. Genau der Security Tradeoff, dass es ein Pfeil irgendwo als Plaintext gibt, was sozusagen vom Agent User besitzt wird. Ist das ein Weg, den man aus Security-Sicht gehen kann?
Oder ist das ein Deal Breaker per se?

  44:53
In DOK, how as I a secret detail for writing secrets, Brandex felt own better agent user.

  44:59
Ja, frage dann in deinem Feedbackbogen, also in deinem Fragebogen. Also, ich würde schon sagen, dass das in dem Sinne in Ordnung ist, weil du es ja entsprechend schreibst, dass nicht jeder lesen darf. Also.

  45:03
Nö.

  45:06
Genau.

  45:11
Versteh ja natürlich, manche Fragen sind hier doppelt, weil sozusagen das eine ist halt für eine ist halt für die quantitative und das andere für die qualitative Analyse. Und da interessiert mich dann immer sozusagen, was die Argumente dafür oder gegen dieses Plaintext wären und wo man da vielleicht Futurework mäßig hinschauen könnte.
Darum steht die Frage da.

  45:28
Also.
gegen also für Plaintex, ich sehe nicht, dass irgendeine Verschlüsselung
Es besser machen würde.

  45:40
Na, das macht's nicht.

  45:40
sage ich mal, du musst jetzt, weil dann muss sie wieder die Cretentials für die Verschlüsselung zwischen der Anwendung, der es lesen muss, und dem Agent teilen. Also wieder Cretentials, die irgendwo pflegen muss und die erst wieder auf derselben Maschine liegen.
Ich glaub nicht, dass man dafür, dass man dadurch viel gewinnen würde.
Wenn ich jetzt nicht plain text wohin schreibe.
Canon.

  46:08
wunderbar. Dankeschön für den Input.
Genau, würde die sozusagen einen weiteren Use Case außer nur Secret Management sehen, was man mit dem Agent abbilden könnte.

  46:28
Konfigurationsinformationen verteilen, also Configuration and Configuration Management. Dieses falsche Wort vielleicht, aber die Frage ist, ob das jetzt.

  46:38
Yeah, Ziel friend is.

  46:39
Ob the office sinnful is pull.

  46:40
Genau, oder ob ich da nicht gleich Kyvinitis nehme, wenn ich ja meine ganze Infrastruktur irgendwie aufsetzen will oder so.

  46:41
Genau.
Not ne, aber not everything is liftable and shiftable. Desto älter, desto.

  46:52
Also, oder, oder, du, du, du meinst OK Konfigurationen, nicht ich stelle mehr Infrastruktur damit her, sondern.

  47:00
Ja, sondern zum Beispiel ja, was andere auch gemeint haben, wär halt dieses Configuration Management Dings, dass dieser Agent dafür verwendet wird, um Konfiguration genauso wie Secrets verteilen zu können.

  47:01
Ich deploy.
Hmm, good I.

  47:12
Never time I an diese Domcat zur Welt gedacht, weil in commonities.

  47:13
bis

  47:19
Ist wieder andere Welt.

  47:19
Also das, das könnt ihr schon sehen, grundsätzlich Property Files, YAML Files, was auch immer, Config Files zu aktualisieren, ohne dass ich direkt auf den Server muss oder komplettes Deployment starten muss, ja.

  47:34
Yep.
Ein Use Case, den ich ein bisschen noch überlegt habe, der vielleicht möglich wäre, wäre, wenn du auf einem Host auch mehrere TLS Files hast, die vielleicht jetzt nicht nur Application relevant sind, sondern h aproxy relevant oder sozusagen Infrastrukturkomponenten sind, könnte man die ja rein kritisch auch abbilden.
Schrägstrich die Credentials vom Agent, wenn man sie überschreibbar macht, könnte man rein so auch abbilden, dass sozusagen ein nur Root Credentials ab hinlegst auf dem Server und dann über dem Agent sagst, alle Fìn alle einmal am Tag, hol dir bitte das neue Certificate von der Asia.

  48:04
Mhm.

  48:11
Und replace dein aktuelles.

  48:11
Mhm.
No.

  48:13
Das heißt, 33 könnte man den Agent dadurch auch ein bisschen sicherer machen, indem er sich sozusagen selber refresht.

  48:20
Mhm.

  48:22
Genau, genau, genau. Seht ihr also, seht ihr außer die Limitations, die ihr schon angesprochen habt, noch andere oder Critical Weaknesses, die Express gehören?

  48:44
Also rein vom.
rein vom vom vom Vortrag jetzt bin ich mal über die Authentifizierung und Autorisierungskonzept, das müssen wir einfach noch mal genauer anschauen. Klar, man kann das alles abbilden, indem in anderen Systemen Dinge so und so und so konfigurieren
Aber gefühlt.
geht jemand her, sagt ja, wenn ASHA ein eine Service Principle oder eine Application App Registration und oder im Key Clock laufen Mois wird keiner hergehen und 200 REMs anlegen, weil ich 200 Produkte in meiner Tomcat Umgebung deploy.

  49:08
Guck mal.

  49:26
Dann habe ich einen und spätestens dann kommen wir darauf, dass das keine gute Idee war, wann jeder die Cretentials von dem anderen überschreiben kann.
Dot, thank you.
Dass man beim Autorisierungs und Authentifizierungskonzept, also zumindest diese Einstellung, nur gewisse User lasse ich rein. Ich entscheide auf meinem Server, welche User da hinkommen und wer nicht hinkommt.
Da scheint mir zu viel Aufwand, das auf der Identity Provider Seite zu machen und da bin ich sehr losgelöst von dem ganzen Thema. Also so eher so ja.
In die Richtung.

  50:05
Versteh.

  50:06
Da bin ich mir noch nicht sicher, ob das zu 100% ausreicht.

  50:06
Common.
Kann man gut in den Limitations erwähnen. Vielen Dank.
Genau, na gut, dann glaube ich, haben wir eigentlich die ganzen Fragen schon abgedeckt. Ich sage auf jeden Fall vielen Dank für eure Zeit.
Sollte euch das Tool interessieren, also ich kann ja mal die Transkription abdrehen.

Transkription beendet
