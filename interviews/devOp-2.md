Transkript
1. April 2026, 12:04PM

 Transkription gestartet

   0:03
Also dann gehen wir einmal rein. Gut, wir haben letzte Woche schon ein bisschen über Secret Management und so weiter geredet und was ich jetzt gemacht habe, ist sozusagen ein hybrider Workflow und einen daraus entstehenden oder gebrauchten Secret Agent, um sozusagen auf einen gleichen Stand zu kommen, das.
Also, weißt du was in Centralized Secret Stories oder soll ich dir das nochmal erklären?

  0:27
Ich gehe davon aus, das ist sowas wie der Secret Server bei uns, oder?

   0:31
Genau, oder der Asia Keyboard oder der Hushikob Vault, es ist ein Harden System, von dem andere Systeme Credentials retrieven können über most likely eine API, die halt den Lifecycle eines Secrets abbildet.

  0:34
Mhm.

   0:47
Genau, gut. Dann sind wir da schon einmal soweit. Der External Secret Operator sagt dir etwas.

  0:57
Ja, den haben wir letztens durchgegangen.

   1:00
In yeah.

  1:01
Es ist die Schnittstelle für die OpenShift im Prinzip zu einem Secret Store.

   1:06
Genau.
Und die Sache ist, ich habe mir das einmal ein bisschen angeschaut in meiner Masterarbeit und propose sozusagen in meiner Masterarbeit einen Secret Management Workflow für Hybrid Environments. Hybrid Environments Consisting out of Linux Servern, die nicht in der Cloud sind, also keine Authentication API haben.
Und Kubernetes Systemen genauso zeigen. Das Mental Model ist, du hast ein Git Repository, wo du für deine Linux Server sowie für deine sowie für deine Kubernetes Sachen deine Secrets eincheckst. NSI ICD Pipeline apply die dann entweder beim External Secret Operator oder
beim Secret Agent auf Linux Hosts. Und die sind halt jeweils dafür verantwortlich, die Secrets abzuholen, abzulegen, als File irgendwo hinzulegen. Und ja, das heißt, das Mental Model, was ich propose als Workflow, ergibt Sinn für dich.

  2:00
Jo.

   2:01
Das gut, genau das heißt, gehen wir es nochmal durch Store CICD Apply, das halt entweder auf Kubernetes oder auf dem Unmanaged Linux Host. Und das Schöne ist, das ist alles Filebased, also in Kubernetes hast du halt deine Koordinative Secrets. Mein Agent erfüllt dann sozusagen den Need.
Diese das, was der macht, auf Linux abzubilden und sie sind beide Vendiagnostik. Das heißt, ich kann relativ viele Secret Stores verwenden. Meine Implementierung kann zur Zeit nur Hashikob Vault und Asia Keyword, aber er ist relativ einfach zu erweitern. Genau.
Dann nochmal kurz XL Secret Operator. Brauchen wir nicht zu tief reingehen. Wir haben ein Azure Keyword oder irgendein Vault. Dann haben wir hier ihn als Proxy, der macht was Komponente Secrets und dann können Applikationen damit arbeiten.
So weit, so klar.

  2:49
Mhm.

   2:49
Und sozusagen dazu brauchen wir dann 2, also der Exile Secret Predator, 2 Mental Models introduced und zwar den Secret Store, der sozusagen eine Konfiguration ist, wie ich mich mit dem Secret Backend connecten kann und zwar entweder
So brauchst du die, welchen Secret Server du ansprichst, also welchen Typen. Dann brauchst du halt den Endpoint und halt die Authentication mit Methode, wie du hinkommst.
Glaub ich ergibt sich.

  3:18
Mhm.

   3:18
Und dann hast du halt ein Secret oder in meinem Fall dann auch genannt Managed File. Das ist halt einfach eine Datei, die sozusagen synchronisiert wird und getemplated wird und generiert wird mit dem Secret Store. Und die ist halt ein Template, wo du im Template eine Referenz auf dem Secret Store angeben kannst.
Und die Location, wo du halt das File speichern willst oder in Kubernetes halt, wie das Secret heißen soll. Und dann gibt es eine Rotation Policy, also wie lange soll das Secret nach Applian gleich bleiben? Das kannst du halt selber einstellen. Und in meinem Fall wären Jason Yammer.
Environment Files und PEM Files überprüft, ob die per se korrekte Files sind, dass du nicht zum Beispiel irgendwas reinschreiben kannst, was sozusagen beim Jason die Formatierung brechen würde über YAML. Genau, und sonst sonst arbeitet der Agent, den ich gebaut hab, halt relativ.

  4:04
Mhm.

   4:09
Sichern schreibt nur 0 600 oder 0644, je nachdem, ob du halt gleiche User oder Gruppe hast, die Dateien.

  4:16
Mhm.

   4:18
Genau. Und demnach meine Idee, die ich jetzt eh schon erklärt habe, ist sozusagen, der XNIC-Corporator ist perfekt für das, was er macht, weil er provided sozusagen eine Vend-Diagnostic Solution. Er ist file-basierend und dadurch über das Filebasierend kannst du halt relativ viel abbilden und
Und er und er hat halt diesen deklarativen Syntax, der halt applied werden kann auf GitHubs Basis. Genau, und demnach fehlt so was eigentlich für Linux-Host, wenn wir ehrlich sind.
Und da habe ich auch Product Research gemacht. Es gibt nichts, was das wirklich abbildet. Es gibt zwar den Haschikop Agent.

  4:47
Mhm.

   4:54
Der sozusagen ist nur für den Vault kann und der spricht halt nicht Jamel, sondern der spricht den HCL, was natürlich auch möglich ist. Aber der hat auch nicht das Konzept, dass du mehrere Secret Stores pro Host haben kannst. So, sondern nur Haschikop Bolz, genau.
Noch einmal den Workflow und das Mental Model, aber das ist dir jetzt eh schon bewusst. So, was macht mein Agent jetzt wirklich und genau? Mein Agent ist sozusagen einfach ein Binary, das irgendwo läuft, eine Rest API exposed.
Und genau diese Restex API erlaubt es dir, Secrets und Secrets Stores zu konfigurieren und er legt die Daten File also in eine SQLI Database, also in ein File. Schreibt er das rein, schreibt er die Connection Strings zum Secret Store rein.
Und halt die Konfiguration, welche Files du gerne getemplated haben wollen würdest. Er unterstützt auch GET, also sozusagen du bekommst auch zurück, was du konfiguriert hast. Bei den Secret Stores kriegst du sozusagen nur die Metadaten zurück und keine authentication Details. Und bei den Secrets kriegst du halt die Konfiguration, die du applied hast, zurück.
Aber nicht die gerandeten Files. Das heißt, die gerandeten Files werden nicht gelesen und verlassen niemals den Host. Genau, was er auch kann, ist halt dieses Reconcilen. Also das, was du ihm gibst, vergleichbar in der Zeit, die du gesagt hast, mit dem Secret Store.
Und bau dir die neue Datei nach der Zeit. Dann haben wir eine Tomic Writer. Dazu kommen später noch, aber der sorgt halt dafür, dass eine Datei wirklich nur replaced wird, wenn's safe ist, weil sonst könntest du ja eine halbe Datei haben und deine Applications acrippeln und das darf nicht passieren.
Genau, und ja, der legt dann halt in einen definierten Folder oder in mehrere. Da gibt es dann eine Sectional Config, auf die ich hinkommen werde. Der hat erlaubte Pfade, wo die Sikrets dann ablegen kann und dort liegen die Sikrets halt dann in Plaintext herum.
Das heißt, dort liegt deine ganze getemplatete Datei als Plaintext irgendwo herum. Genau, und rein theoretisch könnte man ihn hier jetzt in dieser Schematic für halt normale Applications verwenden. Du könntest aber auch Postgres oder andere Filebase Systeme.
Anbinden, indem du zum Beispiel die Pumps oder halt Credentials als Jason hergibst. Oder was er halt auch könnte, ist halt auch Greenfield oder Brownfield oder halt extern gekaufte Systeme zu unterstützen, solange er halt einfach.
Mit den solang das System halt mit den Files arbeiten kann. Ich glaube, das Konzept ist jetzt auch relativ klar und es gibt halt ein DevOps oder eine CICD Pipeline, je nach dem, der diese Konfigurationen halt applying kann, so wie in der Linux Welt, also in der
OpenShift Welt.
Genau. Jetzt kommen wir kurz aus den Features von den Aid vom Agent. Er hat 'ne Crat A. P. I. get wie gesagt Listening Listening Listing von von deinen defined Resources und halt im Fall von einem Store redacted information. Du hast ein Audit Log, das heißt jeder Reroll, jeder Configuration Change
Ist in einer Datei am Host geschrieben. Dann hast du eine die Rerolltime. Du hast einen Slash metrics Endpoint, den du mit Kafana Schrägstrich Promepheu Schrägstrich Alert Manager verbinden kannst, um halt runtime metrics sowie Fehlermeldungen mitzubekommen. Das heißt, wenn er ein File nicht re rollen kann, aus welchem Grund auch immer.
Wird das sozusagen dort hingeschrieben in diesen Endpoint und kann konfiguriert werden, dass es einen Alarm auslöst, wenn du den Agent scrapen lässt. Und die Pfade, also die Root-Pfade, wo Secrets erstellt werden dürfen, müssen explicity explicitly in der Config gelaut werden.
Natürlich könnte man argumentieren, das geht auch einfach nur, wenn der User richtig eingestellt ist, aber per se ist es halt besser, wenn du doppelt sicher einstellst, dass der User nicht Sahn Orte schreiben darf, wo er nicht sollte und das kann man ja vielleicht mal übergeben.
Vergessen oder wenn man den Agent halt ausführt. Als Application User will man ja vielleicht trotzdem nur im Config Folder Sachen schreiben dürfen. Genau.
Und der Agent per se ist halt eigentlich gedacht dafür, dass er im Internet erreichbar ist und demnach brauchen wir auch eine gewisse Authentication.
Weil genau. Und diese gewisse Authentication hätte ich in meinem Agent so abgebildet, dass wir einen MTLS Check haben. Der checkt einmal, ob der Request überhaupt von einem Certified Client kommt. Wenn ja, dann kommt er zum JVT JVT Value Data, der validiert eine JVT.

  9:04
Mhm.

   9:24
Und nur wenn der passt, können wir sozusagen weitergehen. Warum ist das sozusagen eine Dopple Layer Chain of Security? Aus dem einfachen Grund, JVTs wissen, wir haben ein Problem. Und zwar, sobald sie einmal irgendwo geleakt werden, sind sie sozusagen replaybar, bis ihre bis ihre Expiry halt abgelaufen ist.
Weil das halt ein signiertes Dokument ist und wenn die Expirary auf 3 Stunden gesetzt ist, warum auch immer, dann gilt der halt 3 Stunden. Aber dadurch ist es halt stateless. Genau, demnach habe ich ein MTLS Check davor eingebaut, den man konfigurieren kann, um halt doppelte Sicherheit für dieses System zu haben, wenn man es wirklich.
Public Facing im Internet haben möchte genau.

  10:03
Ist es im Prinzip ein kleines Zertifikat oder was? Was ist?

   10:07
MTLS ist einfach Mutal TLS. Das ist ein Standard und das heißt, der Client verschlüsselt, also signiert mit seinem, mit seinem Private Key, wenn mich nicht alles täuscht, den Request.

  10:15
No.
Uh.

   10:19
und oder mit seinem Public, na mit seinem, mit seinem, ich bin mir grad nicht sicher, ob Prior Republic, das müssen wir noch, müsst ich nochmal kurz nachlesen und sozusagen der der Webserver validiert sozusagen den Request, ob der in von den Konfigurierten
Von den konfigurierten Keys, die er hat, signiert wurde. Das heißt, es ist einfach der TLS-Check in beide Richtungen. Das ist standardisiert, das heißt Mutle TLS. Genau. Das heißt, du kannst darüber halt einmal schauen, ob das provided Valid Certificate ist, ob das halt ein Client ist, der dich aufrufen.

  10:39
OK, yeah. Mhm.
Yeah.

   10:52
Darf.
Jessie, da starte ich halt einen Layerden Security Approach.
Um halt das Risiko eines JVTs zu verringern.

  11:05
Mhm.

   11:05
Genau, dann gibt's eine Atomic Files for Protokoll, was ich mir hab einfallen lassen. Als Erstes wird halt der Pfad gecheckt und dadurch kannst du halt keine auch keine Pefter versus machen. Denn selbst wenn du so konfigurierst, dass es Pefter versus wäre, aber du außerhalb eines der konfigurierten Root Verzeichnisse bist.
Kommst kommst nicht weiter, dann hast du eine Preflight Validation. Das heißt, nachdem er die Daten aufgelöst hat, schau da wenn ja.

  11:28
Entschuldigung.
Max, ja, so 22 kurz weg. Ah, die letzte halben Minuten habe ich nicht gehört.

   11:31
Yeah.
Ach so, dann hast du eine Brieflei.
Okay, gut. Du hast einen Security Check, der checkt halt einfach, ob der Pfad da ist, dass du keine Paft Reversals drin hast. Dann hast du so den Content Check, den ich schon erwähnt hab, wo er halt einfach schaut, ist das ein valides Jason, Jaml oder XML oder Pemfile?
Dann hast du ein, dann kreiere ein Temp-File, setzt die Permissions richtig und macht dann einen O.S. Punkt Swap und löscht dann das Temp-File. Das ist halt einfach, um sicherzugehen, dass keine halben Files geritten werden. Demnach ist es Atomic, dieser ganze Prozess.
Hörst du mich eh noch? Passt alles.

  12:11
Ja, ja, ich, ich, es ist noch ein bisschen kleiner. Ich scroll da grad rein, aber aber ich glaub inhaltlich hab ich schon verstanden.

   12:16
Ach so, ja.
Ja, ist du musst es nicht. Also, du, du musst ja keine Prüfung überschreiben. Es geht sozusagen, dass du einfach siehst, dieser es wurde an sozusagen Atomic Swaps gedacht. Genau, dann Installation Installation of the Agent. Es gibt ein Playbook per se.

  12:26
Ja.
Ja, das ist gut.

   12:35
Das Playbook ist so konfiguriert und könnte Snippets auch rausgenommen werden, wenn man es anders machen wollen würde. Und zwar, es kreiert den System User, es kreiert eine Secret Group. Also in der Group sollten dann alle alle Accounts drin sein, die das lesen dürfen.
Es erstellt die Pfade, die halt der Agent braucht. Also weißt eh so wie in Linux, wenn du ein Binary installierst, hat das halt immer gewisse Pfade für Logs und so weiter.
Genau, it's Templated dir die Konfiguration, die wir uns dann gleich anschauen werden und es created dir ein System D-Service, dass der Agent halt läuft und halt auch läuft und einmal restarten muss und so weiter. Und es lädt das Binary halt einfach runter, das ist sozusagen.
das Install, das Install Behavior und ich finde, das ist relativ gut vergleichbar mit einem mit einem Install für für Kubernetes, weil wenn du dort den Agent installieren willst, müsstest du es ja auch mit Cube CTL oder mit irgendetwas machen.
Und so hast du es rein theoretisch auch Hidden Away in einem pre-defined Weg und du kannst halt mit Variablen das Behavior verändern. Genau, was kann der Agent jetzt genau konfigurationsmäßig? Du kannst ihm halt einen Namen geben. Das ist halt wichtig für Scrapen dann.
Dann, du definierst sozusagen alle Roots, auf die er hin darf. Das heißt, das wäre in unserem Fall vielleicht slash www.slash Tomcat MB slash.com sozusagen schön herein reingecoded, wo er genau hinschreiben darf, weil sozusagen, du willst vielleicht nicht, dass er auf den Logs.
Dings irgendein Fall schreiben darf und so weiter. Das heißt, das kannst du hier sozusagen sehr fein granulat steuern. Das heißt, selbst wenn du den gleichen User verwendest für mehrere Sachen, was du nicht solltest, hast du hier einen kleinen Security Mechanismus eingebaut. Dann kannst du dich beim Log Level halt zwischen Info und Jason wählen und wo du hin.
Du gern das Audit Logfile geschrieben hättest, dann musst du angeben, wo die lokale SQLite Database hingehen soll. Das Schöne ist der Agent Package, das SQLite Binary. Das heißt, du boah, der das heißt, es gibt keine Dependency auf SQLite am Host.
Und security mäßig, wie gesagt, kann man TLS aufdrehen und dann kannst du halt hier das Client Certificate File angeben und dadurch kannst du halt dann MTLS machen und du brauchst halt auch ein Cert und ein Key File für TLS, dass du TLS anbieten kannst, aber das sollte jedem bewusst sein.
Das ist sozusagen auch das größte Problem des Agents und das ganze Workflows noch, aber das hast du auch beim External Sequel Operator und zwar, dass du sozusagen Root Credentials irgendwo brauchst. In irgendeinem Verfahren musst du ihm halt sozusagen eine gewisse Art an Credentials mitgeben.
Weil das sozusagen zum Bootstrappen braucht.
Genau, dann können wir uns bei den Authentication Modes zwischen Non M.T.L.S. und J.V.T. entscheiden. Jetzt, wenn du jetzt wirst dich fragen, warum kann ich das auf M.T.L.S. setzen? Das Schöne ist, du könntest M.T.L.S. auch verwenden als Authentication. Das heißt, du hättest halt nur den M.T.L.S. Check.
Aber der secureste Weg, den ich empfehlen würde, wenn man im Internet hängt, wäre ein MTLS-Check, gefoldert von einem JVT-Check. Und im besten Fall hast du einen Identity Provider und könntest halt hier angeben .wellknown.jvcase.json, das ist sozusagen
An eine D.P. gibt es sozusagen einen Ort, wo sich Clients Public Keys abholen können at Runtime, um einen G. V. T. zu validieren. Genau, und Issue und Audience sind einfach auf Sachen, die man eintragen kann.
Muss man nicht, hilft einfach nur, macht das Ganze noch ein bisschen besser von Audit Lock. Und genau, du könntest auch reintutschen, Shared Secret für den JVT haben oder halt selber da den Public Key reinschreiben. Aber das ist halt nicht so secure, es ist und so einfach.
Also, meine Empfehlung wäre, einen IDP zu verwenden für kurz gelifte JVTs und halt MTLS genau.
Gut, kommen wir auch schon zu Configuration Examples. Du siehst hier auf der Seite den Asia für beide Systeme eine Konfiguration des Asia Stores. Genau, das ist der externe Secret Operator. Die Konfiguration, das ist die Konfiguration für das System, was ich gebaut hab für die Linux Hosts.
Er spricht nicht 1 zu 1 die gleiche Sprache, weil ich das persönlich nicht so gut gefunden habe. Aber was deine Meinung dazu ist, werde ich dann notieren und dann vielleicht und dann sozusagen erwähnen im Future Work Part meiner Arbeit. Genau. Also, was wir hier sehen, die externen CCorporator, die haben es sozusagen ja vorgemacht und mir vorgedacht.
Bei denen sagt man hier den Typen, dann sagt man sozusagen die Tenant, dann sagt man Tenant ID. Die URL finden wir auch hier und hier und das Client ID und die Client Secret. Wir sehen hier in Kubanetis, ziehen Sie sich das halt aus einem anderen Secret. Ich bin den Weg gegangen und
Ich hab gesagt, okay, wir sind dort auf einem Plattform ohne diesen ganzen Secret Konzept, was Kubanetische nativ hat. Demnach schreiben wir das gleich mal hier rein. Hier sind die Werte natürlich redected, weil ich glaub, ich brauch keinen kleinen Secret Antennen ID zeigen. Das schaffen die Leute schon selber.
Genau.
Gut, kommen wir zum Configuration Example für einen von einem Secret. Wir sehen hier wieder das vom externen Secret Operator. Der hat hier die Referenz zu dem konfigurierten Vault, also zu dem hier, zu dem hier mit dem Reference Name, genau.
Das heißt hier, das musst du halt angeben, von wo er sich zu welchem Wort er sich verbinden soll, pro templated and secret. Und dann schreibst du halt hier ein Template. Und das ist halt dieser Go Templating Language, schreibst du dann dort. Dieser Wert muss mit dem Wert übereinstimmen, mit diesem Key.
Und dann kannst du eben hier beim beim Exile Secret Operator noch mal sagen, dass da remote Reference dann so heißt. Und daraus baut er dir dann das Secret mit dem Namen TCN Secret in OpenShift/Kubernetes. Das ist dir klar?
Wie das einmal funktioniert, konzeptionell.

  18:35
Mhm, das ist jetzt Standard External Secret Operator oder?

   18:38
Genau das ist Standard Excel Secret Operator Definition. Genau, und auf der anderen Seite sehen wir, wie ich das abgebildet habe.

  18:41
Mhm.

   18:47
Ich hätte das so abgebildet, dass mein Ding auch einen Namen hat, dass es halt einen File Peff hat, weil im Vergleich zu Kubernetes kannst du nicht nur mit Namen arbeiten, sondern brauchst irgendwo einen Pfad für die Datei, dann einen Store Name, so wie da auch, weil du eigentlich hin verbinden musst.
Dann ein Refresh Interroll. Ich hab's Virol Time genannt. Genau, und dann gibt's halt.

  19:09
Was macht der Refresh-Intervall genau? Prüft der jetzt alle?

   19:11
Nach 25 Sekunden baut er das Secret neu. Also nach 25 Sekunden holt er sich die Daten neu vom Secret Store, legt die neue Datei ab und swappt sie dir.

  19:25
Okay.

   19:26
Genau, und bei mir hast du dann halt hier in Secret Rapper den Content, dann hier das gleiche Example und ich hab das halt ein bisschen anders gelöst, weil ich ich find das persönlich schlechter. Aber da ist mir auch deine Meinung, kannst mir dann gerne deine Meinung sagen. Ich hab das so gemacht, dass du sozusagen hier
Das Match, das und so heißt in der, so heißt sozusagen im Remote.
Es ist halt einfach, finde ich, ein bisschen weniger Verbots, weil du sozusagen hier das Match du so heißt es hier, so heißt es im Secret Store genau.
Aber sozusagen konzeptionell ist dir das auch klar, wie das jetzt funktioniert. Das ist wichtig, weil du kriegst gleich eine Aufgabe dazu.

  20:11
Uh.

   20:14
Ja, jetzt kriegst du sozusagen, dass das nicht nur ein Konzept ist, sondern dass du mal auch ein Secret configured hast. Schicke ich dir jetzt sozusagen das Secret, was ich dir gerade gezeigt habe, als Text und bitte dich dann darum, mit den Anforderungen, die ich dir da gerade gegeben habe, dieses Secret zu bauen, dass du halt einfach.
Einfach kurz ein bisschen Erfahrung hast und siehst, wie das Ding funktioniert.
So.
Uh.
Oh nein.

  20:42
Aber du hast doch gesagt, ich muss keinen Test schreiben.

   20:45
Es ist ja kein Test, es ist ja nur ein Look and Feel.

  20:48
Ah, OK.

   20:49
Sure.
Und das nicht schaffst.

  20:51
Muss ich mir merken, das muss ich, das muss ich meinem Buben auch mal sagen, wenn er, wenn er Schularbeit hat.

   20:57
Ich will ja nur wissen, ob du es könntest.

  21:00
Ja.

   21:00
Genau, Secret Store soll Dev Store sein. Die Reval Time soll eine Stunde sein. Der Content soll sein Secret und den Variablennamen, den du geben möchtest. Und dieser Variablenname soll auf den Store Key Secret Dev referenzieren. Pfeilpfad und Name sind irrelevant für diese Aufgabe.
Genau.

  21:22
Das heißt, du willst jetzt, dass ich das Text-Snippet anpasse oder wie oder was?

   21:25
Genau, genau. Du kopierst sie dir in ein, du kopierst sie dir einfach in ein Notepad, editierst das kurz, dass du halt das auch einmal selber gemacht hast.
Weil ich, ich finde sozusagen, du hast dadurch halt ein bisschen mehr Verständnis von dem Ganzen, wenn du es schon einmal selber ausprobieren müssen.
Dann schickst du mir das wieder, dann schau ich mir das kurz an.

  22:05
This is done a.
Jason oder was?

   22:08
Ja, in dem Fall ist es ein Jason, genau.
Einfach, dass es nicht das Gleiche ist.
Kommst du gut fahren? Hast du eine Frage?

  22:54
Na, ich probier da ein bisschen mal. ***, schick es dir gleich zurück und und der Pfeilpfad ist das dann der gleiche oder ist das Wurscht oder?

   22:56
Yo, yo, du.
Pfeilpfad oder Name sind egal für die Aufgabe.
Also das heißt, die kannst du unverändert vom Template lassen.

  23:19
Ja, du hast ja XML geschrieben. Da kann ich kein XML draus machen.

   23:23
Mhm.

  23:23
Es gibt.

   23:27
Können schon.
Das ist dem Agent egal. Du kannst den X. L. nennen und drinnen steht, wobei er wirds dir nicht erlauben abzuschicken. Also, er wird sie nicht speichern, aber per se. Er wirds einmal auflösen probieren.

  23:34
Yeah, yeah.
da
Unterstützt der das mit 1 H oder oder muss man das mit 3600 schreiben? Passt du das?

   23:50
Ja, er unterstützt einfach. Er unterstützt bis.
Ich glaub bis Day sogar.

  24:03
Jetzt muss ich aber schauen, ob da einrücken, ja.
Genau, die Secret Reference ist im Secret Wrapper drinnen. Habe ich das schon richtig verstanden?

   24:18
Trade refer.

  24:19
Ja.
Content on secret references on of an a ebene.

   24:23
Yeah, you know, yeah.
Das hab ich schlecht notiert geschickt.

  24:30
Das ist gar nicht mal die Geschichte.

   24:33
Yeah.

  24:35
Don't met this super says.

   24:35
You are this is teams.
Ah, Schwupp.
Ja, wunderbar. Du hast es geschafft, das zu konfigurieren. Secret auch, Secret Dev und Rewall Time eine Stunde. Secure Location sind wir, haben wir sogar. Bum, speichern wir speicherns in Secure Locations, oh.

  24:48
Juhu.
Ja, natürlich, mit Jason sogar magst du.

   24:59
Kommen wir schon zu der zu dem Form, was du jetzt ausfüllen darfst. Die Fragen, wenn dir irgendwas unklar sind, frag mich einfach. Darum ist das per se ein semistructured Interview, dass das sozusagen und nicht nur eine Umfrage.
um halt sozusagen das ja machen zu können. Genau danach gibt's dann noch ein paar Open Questions, die du nicht alle beantworten musst, aber alle Fragen und so weiter, die du hast basierend darauf, kannst du mir dann gerne stellen und am Ende muss ich dann noch ein
Also, nach der Serie stelle ich dir kurz noch ein paar andere Fragen, die ich vergessen habe, die einfach nur zu dir und zu deiner Person sind.
Aber ja.

  25:45
Genau jetzt, jetzt hab ich aber schon bei der ersten Folge schon einmal eine Frage. Magst du das Chart noch mal kurz herzeigen?

   25:49
No.
welches

  25:53
Mit mit die Limited 3 Linux Server, was du da gehabt hast und dem dem der CICD genau.

   25:58
So.
No.

  26:06
Was macht jetzt die CICD ganz oben? Das habe ich jetzt glaube ich nicht.

   26:08
Die CICD nimmt die Datei, die du geschrieben hast, gerade und applied sie.

  26:13
Ja.
Das heißt, das ist im Prinzip die Konfiguration für deinen Client, damit er weiß, was er machen soll.

   26:21
Genau.
Und so, so wie das jetzt hier Grad mit Helm machst, würde die CICD halt auf den Host, dann halt die Dateien über die API draufspielen.

  26:23
On.
OK, ja, passt und die Authentifizierung.
Du hast ja geschrieben, genau.
Der Client, was ist jetzt da herinnen, der Client und der Server?

   26:54
Das der der Server ist der Agent, der auf einem Host rennt, und der Client kann entweder du sein, ein DevOps, der mit Curl arbeitet, oder eine CICD Pipeline, die auch mit Curl arbeitet.

  26:59
Genau.
Ja.
Was macht der Client mit dem Agent

   27:10
Der schickt ein HTTP Request.

  27:16
Also mein Verständnis ist der Client, der läuft auf dem Linux Server oder der Haut die Konfiguration.

   27:16
der
Ja.
Nein, nein, nein, nein, nein, nein, der Client bist du. Das ist eine CICD Pipeline, das ist der Jenkins, genau der Server, der Teil, das ist.

  27:25
Ah, the Entschuldigung, the server, server money heads.
Läuft am Linux-Server, der hat eine Konfiguration von der C.I.C.D. Der prüft alle X. Sekunden, ob die Secret noch, also der tut es updaten, mehr oder weniger, mit dem Atomic Change.

   27:42
You know.

  27:45
Der, du musst jetzt einmal mit Prometheus und Graphana vielleicht erreichbar sein, aber aber was, aber für was brauche ich denn den Rest Schnittstelle, das, das, das habe ich jetzt glaube ich nicht verstanden.

   27:52
Was?
Für was brauchst du die Restschnittstelle? Na ja, irgendwo müssen ja die Store Configs und die Secrets, die pro Host da sein sollen, konfiguriert werden.
But when.

  28:06
Genau, aber das macht jetzt die, die die CISC oder oder greift die CIS. Also OK, das heißt für die für das Ausrollen der Configs brauche ich die Schnittstelle, das heißt das.

   28:16
Genau.
Genau, einfach ein.

  28:24
Und du hast erwähnt, der muss im Internet verfügbar sein.

   28:28
Es kann, es kommt halt, es, er ist, er sollte sicher genug sein. Du hast dann gleich eine Frage dazu. Wie du ihn bei dir in deinem Wireless aufsetzt, ist komplett egal. Das sollte halt einfach nur sicher genug sein, dass du ihn per se ins Internet hängen könntest.

  28:35
OK.

   28:44
Wenn du das möchtest.

  28:44
O. K., das ist das Ziel. Er muss jetzt per se nicht im Internet sein, das heißt, das ist einfach irgendein Port, der am Linux-Server halt erreichbar sein muss und die C. I. C. D. muss auf die, das wäre zum Beispiel bei uns fast sogar ein bisschen problematisch, aber du kennst das.

   28:48
Ne.
Genau.
Genau, und du könntest auch ein.
Ich, Ich weiß das, aber per se ja per per se könntest du ja halt rein natürlich auch ein Skript am Server haben, was dir den, was dir dann diesen, was dir den Curl Aufruf macht.

  29:01
Thema vermutliche.

   29:10
Du könntest ja auch den den Dings nur für Localhost exposen in unserer Umgebung und die C.I.C.D. triggert halt das. Per se ist das natürlich dann 1, weil weil eine spezielles Environment so ist, aber sozusagen denk drüber nach, wär es, wenn du sozusagen ein Public Facing also.

  29:13
Ach so, du. Ja, ja, genau.
Ja.

   29:29
Würdest du dich trauen, das Ding Public Facing so mit diesen zwei Security Steps zu deployen oder nicht? Das ist sozusagen einfach dann die Frage. Das gibst du halt einfach deine subjektive Antwort aufgrund deines Expertenstatuses und damit.
Werde ich irgendwas machen.

  29:46
Ja, ich, ich gehe mal die Frage.

   29:47
Jo, hello.

  30:09
Der Agent ist go, oder hast du gesagt?

   30:12
Der Agent ist in GO geschrieben, genau.

  31:33
Was ist damit gemeint? The separation of secret store connections and secrets?

   31:38
Yeah.

  31:39
Drift Secrets Store connection is the store. The Secrets and die Secrets und die Separation. Was ist mit der Separation Bond?

   31:43
Yeah, genau. What is the genau?
sozusagen, dass du in den Secrets keine nicht die Details der Connections zum Store brauchst. Ob du findest, dass sozusagen dieses Teil in es gibt Secret Store Configuration und es gibt Secret Configurations, die sozusagen die Store Configurations nur per Name.
Verwenden, ob du das gut findest und ob du meinst, ob das sich.

  32:09
Ach so, weil du in der Config nur DevStore referenzierst, aber nicht mehr.

   32:13
Genau.
Genau.
Und was du sozusagen natürlich auch kannst, du hättest könntest eigentlich ein Repository haben, wo du die ganzen Store Configuration managed, managed, wo sie sozusagen nur DevOps hin dürfen oder nur privilegiertere Leute und 1 haben, wo sich Entwickler sozusagen selber ihre Configuration und Secrets warten können, ohne dass sie jemals
Irgendwelche Connection-Details vom Secret Store sehen müssen.

  32:40
Mhm.
Das heißt aber dann für einen für einen Agent gibt es dann im Prinzip zusätzlich auch noch eine Secret Store Configuration, oder?

   32:54
Genau, der Agent kriegt die Store Configuration und die Secret Configuration. Die Store Configuration haben wir uns ja hier angeschaut beim Actional Secret Operator und beim Azure Keyboard. Das heißt, du musst ihm das schon natürlich geben, wie er hin darf und das liegt in der SQLite Database.

  33:08
Mhm.

   33:10
Vor Ort dort am Client herum und dort liegt es halt im Secrets-Bereich von OpenShift herum.
Und mit den Daten connected er sich halt hin. Das ist halt sozusagen eine Lösung, wenn du keine Managed Identities von dem Host hast und die meisten Systeme, die du sozusagen betreibst heutzutage haben noch nicht Managed Identities.
Solange du also insofern du nicht ja.

  33:32
Mhm, ja.
Das heißt, das liegt dann in einer SQLI Database drinnen. Wie kriegt das der Agent überhaupt?

   33:39
Genau.
Na ja, du appliest es ja, du schickst das als HTTP Post Request hin.
Und dann ist es sozusagen verfügbar in seiner Datenbank. Das heißt, du schickst genau, genau das geht alles überall, alles über REST, so wie beim Maxi Channel Secret Operator.

  33:51
Das wird nicht als File abgeklickt, sondern das geht nur über Rest.

   34:00
Du schmeißt ihm halt über Rest die Sachen so hin, wie du möchtest, und er arbeitet dann damit.
Alles klar bei dir?
Mhm.
Hello.

  37:06
Pers di.

   37:07
Warst du absichtlich draußen oder unabsichtlich oder?

  37:10
Nein, mein Internet hat entschieden, dass.

   37:11
Ach so, überhaupt kein Problem. Ich hab so, ich hab grad mal so gefragt, ja, wie geht's dir so? Hab keine Antwort bekommen, war so. Ich werd auch keine Antwort bekommen, wenn du nicht da bist.

  37:13
Ausmachen will, sorry.
Ich hab nix mehr gehört mit.
Hier rede ich nicht mehr.

   37:27
Hä, so eine blöde Aufgabe. Jetzt hast du mir 15 Minuten was erzählt und jetzt muss ich da deppert Kreuzerübungen machen.

  37:34
Genau, nein, nein, es ist halt jetzt nichts mit Jetzter, warte mal.

   37:36
Ja, ja.

  37:40
Ja, lustigerweise.

   37:42
No.

  37:42
Ich habe nur Probleme, wenn ich in irgendwelchen Depperten Teams bin. Keine Ahnung, woher das kommt.

   37:48
Ja, es liegt es liegt im IQ der anderen Person, wenn der Halt niedriger ist als deiner, dann will dein Internet sich das nimmer anhören mìssen.

  37:55
Was?

   37:56
No, no, no.

  38:01
Nein, das weiß ich nicht.

   38:03
Ja, ja, nein, Spaß, also jetzt nicht mehr wirklich sehen.

  38:05
So, so, what do my doing was?

   38:08
Hört sich aber nach wirklich sehr interessanten Problemen.

  38:12
Ja, keine Ahnung, das Modem hat es jetzt gerade irgendwie resettet. Ich weiß nicht, ob das irgendein Routing-Thema ist oder sonst irgendwas, den Provider anrufen.

   38:17
Yeah.
Ja, ja.

  38:21
Tut mir leid.

   38:23
Macht ja gar nichts.
Solange ich meine Antworten von dir bekomme.

  38:29
So let me this quick with the week.
Da der Händling auf Secret Vice.
Sure.
Man zur letzten Security Frage, wo du schreibst, die Agents Authentication Architecture is sufficient to allowing Internet basing exposure.

   39:12
Ja.

  39:15
Internet Facing Exposure, da ist halt immer mehr dabei. In unserem Fall würde man nie einen Linux-Server ins Internet stellen, in keinster Architektur.

   39:21
Ja, ja, du, die Sache ist, das geht auch viel mehr als deine Experten in deiner Expertenmeinung. Würdest du sagen, du würdest dich konfident fühlen, security mäßig den auf einer V. M., wenn das eine Public Host V. M. ist, die erreichbar ist?

  39:28
Yeah.

   39:37
Zu Alone oder nicht?
Ist ja auch okay, wenn du nicht der Meinung bist.

  39:44
Kann ich über die API Secrets abrufen? Net oder?

   39:47
Nein, du kriegst nur die Konfiguration, die du hingespielt hast.

  39:51
Ah.

   39:53
Also, du kriegst sozusagen die Liste und Konfigurationen. Du kriegst, welche Secrets du aus der definiert hast, aber die Authentication Details sind alle redacted. Die kriegst du nicht vom System runter.

  40:03
Umm.

   40:11
Wenn du die wirklich brauchst, kannst du dich im SQLite Client am Host mit der Datenbank auseinandersetzen.

  40:21
Stehen dann dort die Credentials drinnen, also die Passwörter, auch die Secrets aus dem Secrets Store?

   40:23
Ja.
Die Secrets zum Secrets Store? Nein, die also die Secrets liegen nur in den nur die Connection Details zum Secrets Store liegen dort herum.

  40:31
die nicht nur die
Warum, warum hast du dich dich dafür SQLite entschieden?

   40:41
Punkt 1, weil ich es einfach einpacken kann in den Binary und ich demnach keine Dependency auf irgendeinem System Dings hab. Und weil sobald sozusagen der Host kompromittiert ist und jemand mit Root Access File Access hat, auch die ganzen Secret Values sowieso exposed sind.

  40:47
No.
Ja, schauen wir mal, aber was, was wäre jetzt alternativ, dass du sagst du, dass du irgendwo in einem sicheren Bereich jason, weil jeder kann ja das Collette man kann aufmachen.

   41:06
Ja, es collided einfach aus dem Grund, weil es halt einfach programmiermäßig leicht und schön zum Ansprechen ist und du halt damit mit mit mit Time Queries arbeiten kannst, um halt die Revolts abzubilden.
Das ist sozusagen der größte Grund für Esculate.
Um sozusagen einfach die ganze Arbeit.

  41:25
Also die interne Mechanik im Prinzip, was der Agent, OK.

   41:28
You know, yeah, yeah.
Und du hast halt auch den Vorteil, dass halt dieses Insert und Delete und so weiter standardisiert und gut funktioniert.
Weil nachdem wir sozusagen die Storekonfigs per REST API und die anderen Konfigs auch per REST API in das System schießt. Wäre es halt file-basiert dann halt wieder ohne halt Database Engine halt wieder sehr anstrengend, weil dann musst du halt im File trunkaten, verändern und so weiter.

  42:11
Mhm.
Genau, Hybrid Workflow oder generell Hybrid Szenario in deinem Fall ist einfach nur Linux die Agent und Open Shift ist halt dann der Operator, den es halt eh schon gibt.

   43:15
Yeah.
Genau, genau. Weil sozusagen diese Konzeptional, also die konzeptionelle Welt, sozusagen dieses Store und Secret Configuration auf beides anwenden kannst.

  43:35
Mhm.
The separation between secret authoring and secret delivery is clear and beneficial. Was meinst du mit authoring und delivery?

   43:52
Naja, du schreibst ja per se nur die Konfiguration, wie das Secret ausschauen soll, aber die Delivery des Secrets macht ja sozusagen nur noch der Agent. Das heißt, keiner sieht mehr die Secret Values, außer die Leute, die auf dem Server Rechte haben, diese Dateien anzuschauen.

  44:13
Mhm.

   44:14
Weil dadurch, dass du halt sozusagen nur noch in beiden Fällen diese Templates hast.
Könntest du schreibst du das halt?
Und würdest niemals sehen, was DB User, DB Passwort und DBURL ist, wenn du nicht die Berechtigung dazu hast, entweder in OpenShift diese Names bei Secret anzuschauen oder in Linux halt die Userrechte.

  44:51
Genau, und die zwei Files, die du gerade offen hast, sind ja dann im Prinzip Files, die von einem Applikationsteam gewartet werden, weil die sind ja.

   44:58
Genau die oder vom DevOps oder von wem auch immer das möchte, aber es könnte auch ein Developer machen, ohne jemals ein Secret sehen zu müssen.

  45:01
Genau, ja, mhm.
Genau.
Genau.
Was ist E.S.O. bei der Frage, Linux Secret Action ist Concept?

   45:12
External secret, external secret operator, hab ich mich wahrscheinlich verschrieben.

  45:18
Nein, nein, externes, das macht Sinn. Was meinst du mit Well Align im Sinne von, dass der kompatibel ist?

   45:23
Konzept.
Conceptually, weil die sind ja sozusagen nicht 1 zu 1 gleich, ob sie einfach konceptually, ob du meinst, ja, das passt zusammen. Das würdest du als als guten Workflow sehen, als konzeptionellen.

  45:27
Ach so, von den gleichen Konten, ja.
Secret Rotation wäre dann direkt eine Frage vom Azure Secret Wall zum Beispiel. Ist jetzt aber nicht zwingend Teil von deiner Arbeit.

   46:44
Per se rotieren würdest also per se rotieren würdest du natürlich, weil der Lifecycle des Secret Werts im Wald ist, im Wald.

  46:45
Also, wie man es macht.
Mhm.

   46:54
Und delivered wird es dann halt über die Rerall Time werden oder indem du das Secret neu appliest.
Also, wenn du jetzt einen Breach hättest und es halt konfiguriert ist, dass es nur alle 12 Stunden updated, wäre das ja blöd, dass du jetzt 12 Stunden mit wenigen Redentials oder mit geleakten Redentials arbeitest. Demnach könntest du sozusagen dann in deiner CICD Pipeline das Verfahren starten, einfach.

  47:00
Mhm.
Mhm.

   47:16
Alle Secrets noch mal zu applyen oder halt die ausgewählten mit den Werten. Genau, aber das ist ja.

  47:21
Du musst, Du musst aber glaube ich auch mit bedenken, dass du dann zwar das das X. M. L. oder das JSON oder was auch immer updatest, aber die Applikation das zwingend noch nicht mitkriegt, weil sie zum Beispiel erst restarted werden muss.

   47:29
Yeah.
Das ist halt Application Dependent. Das ist ja Application Dependent. Manche Applikationen schaffen es alleine, andere Applikationen, also manche Applikationen watchen ja die Files, die sie als Konfiguration haben. Aber das ist einfach Applikation, die dependent demnach.

  47:45
Mhm.
Aber das, das sagst du, das ist nicht Teil von von deinen.

   47:51
None.

  47:53
Post.

   47:54
Man könnte als sozusagen dann bei den offenen Fragen als fehlende Sachen reintisch mitnehmen, dass es ein ein Event oder ein Signal gibt in Linux.
Aber man könnte doch einfach sagen, man kann auch einfach die Applikation so umbauen, dass die halt einfach das die Datei watched, weil wenn ich eine Rerolltime, also wenn ich möchte, dass mein Secret rerollt wird, dann möchte ich auch, dass das funktioniert per se.

  48:07
Ja, das ist.
Ja, ist halt, also war natürlich am schönsten, wenn es die Applikation direkt kann. Solche Reset-Szenarien sind halt relativ komplex, weil du wirst ja dann nicht sofort alles restarten und dann ein Downtime haben und und.

   48:31
Ja, die Sache ist, das ist die Sache ist jedes, also jede Firma schriftlich, jede Person verwaltet ihre Applikationen ja anders.
Das heißt, es gibt nicht wirklich einen schlauen Weg, wie du in dem Agent einbauen kannst, wie eine Applikation restartet.

  48:47
Eben, eben, ja.

   48:48
Du könntest einbauen, dass du, dass du kalt konfigurieren kannst, ob es halt System, die Sachen sind und es wird halt einfach massiv komplex. Ich glaub, das ist einfacher und besser, wenn das Problem sich jeder selber löst.
Wie gesagt, mit einem Filewatcher. Passierend darauf, aber ja, das ist ein Punkt, den man mitnehmen kann. Aber das der Ziel, das Ziel des Agents ist halt, dass er nicht, dass er nicht in die Runtime, also dass er, dass wenn der Agent down ist, dass Applikationen trotzdem noch arbeiten können.

  49:11
Ja, ich glaube auch.

   49:21
Das ist halt das Ziel des Agents. Das heißt, der Agent will nichts. Also, dieser Fileroller will nichts damit zu tun haben, wie deine Applikation gestartet wird. Der will nur Files ablegen.
So bist du schon bei der letzten Kategorie.

  50:23
Ja, die schauen wir uns noch mal ganz kurz durch.

   50:25
Nein, nein, ich frag nur, weil für die Frage, weil ich es dir nicht gezeigt hab, sie ist easy to integrate, um sozusagen mit dem Agent sprechen zu können, musst du nur ein Curl, dann halt den Typ Post oder Delete oder Get und dann kannst du minus minus binary Data und halt.

  50:36
Ja.

   50:41
dann die Datei reinlinken. Das heißt,

  50:44
Genau, also davon wäre ich ausgegangen, wenn du sagst, der Rest Call, dass du nicht dort das rein verpackst. Das einzige, was halt ist, mit in dem Rest Call müssen halt dann im Prinzip die Secrets vom Secrets Store drinnen sein, oder?

   50:46
Wunderbar.
Yep.
In welchem von von der C vom C von der store configuration, yeah.

  51:00
von der C. I. C. D. Genau richtig. Wäre in unserem Fall dann eher blöd, weil wir sagen, weil speziell bei uns, vielleicht auch bei anderen Firmen, die C. I. C. D., sprich Jenkins zum Beispiel, relativ offenes System ist, weil es halt eher zum

   51:14
Yep.

  51:18
Bauen von Artefakten verwendet wird, wo relativ viele Entwickler drauf sind und wir zum Beispiel sagen, wir werden da, wir wollen da bewusst keine Secrets oder so wenig Secrets wie möglich drinnen haben, sprich.

   51:28
Mhm, der also per se, der der Jenkins oder was auch immer das Applied braucht halt die Credentials zum zum Agent, aber das hast du auch bei OpenShift und so weiter und bei OpenShift so. Die dieses Ding, was

  51:37
Genau.
Genau, es ist halt die Frage, ja, es ist halt die Frage, ob wir den den external Operator auch über Jenkins aufsetzen wollen oder ob wir uns nicht da einen separaten Prozess, so wie es wir jetzt zum Beispiel bei die Decrets lokal ausmachen,

   51:45
Thus thing was a blood.
Ja.
Mit 10 DB, ja, ja, ja.

  51:58
Eben genau aus dem Grund, weil theoretisch jeder, der den Job sieht, kann im Replay nachschauen, was war da drinnen, kannst dir vielleicht Sachen auslesen, was ja bei so einem Secret Store natürlich ziemlich fatal wäre.

   52:01
Genau, aber ja.
Ja genau, aber da ist es schön, dass du sozusagen Secret Stores von den Secrets trennen kannst und sozusagen dann nur noch.

  52:11
Mhm.
Das ist super. Ja, das ist auf jeden Fall super.

   52:18
Deine paar Connections managen musst und das Schöne ist, du kannst halt dann per se auch einfach diese Connection updaten und alle Secrets, die auf dem Host die Connection verwenden, sobald sie re rollen, würden halt die neuen Daten hernehmen.
Und das ist sozusagen auch einer der großen Vorteile gegenüber von Secrets abholen lassen von der Applikation, weil du hast Punkt 1 keine Runtime Dependency auf das und du brauchst halt nicht die Zugangsdaten zum Vault der Applikation geben.

  52:43
Mhm.
Yeah.

   52:47
Weil natürlich, dass du dir als Applikation was holen kannst, brauchst du halt erst recht wieder die Credentials. Und die kriegst du halt entweder über so Managed Identities oder halt über so wie beim Secret Store. Das heißt, irgendwo muss die Applikation ja irgendwas haben.
Um ihre Identität prüfen zu können.
Das heißt, selbst wenn du das einbaust in das Abholen und der Blue Code für dich cool ist und du das gerne für 17 Systeme machen möchtest, dann nein, nein, ich führ nur kurz den Punkt weiter aus für das Ganze. Dann müsstest du halt auch alle 17 Systeme neu starten, wenn sich irgendwas mit deiner Secret Store Connection ändert.

  53:12
Tana.
Ja.

   53:22
Was du halt hier auch nicht hast, sondern hier ist es halt gekappelt voneinander.

  53:25
Genau, das find ich gut. Was ich noch sagen wollte ist, du kannst daher gehen, du musst das ja nicht, ich denk jetzt nur, wie gesagt, bisschen in unsere Arbeit. Du musst das ja nicht von der C.I.C.D., sondern du könntest das ja auch von einem DevOps-Rechner oder von irgendeinem Rechner, der halt die Rechte hat, damit es dann auch wiederum in keinem in keinem System drinnen ist. Genau, du musst halt dich authentifizieren können.

   53:31
Ja, gerne.
Ja, genau.
Solange du den JVT und den J.
Ja.

  53:45
Aber das kann man ja eventuell irgendwie einrichten.

   53:47
Ja, solange du genau.

  53:49
Unterstützt der mehrere Credentials? Der Client, also der Agent, dass ich sage, ich habe.

   53:53
Remends.
Du könntest, du könntest 5 Stores, du könntest 30 Stores haben pro pro klein, das heißt du ja.

  54:00
Ja, schon, aber ich meine jetzt eher das Client-Zertifikat. Man bei Client-Zertifikat wahrscheinlich weniger, aber Chart.

   54:06
Klein, klein, zertif chot mäßig. Chot ist ja das Schöne. Warte, gehen wir noch mal auf die Konfiguration. Chot mäßig unterstützt er entweder du gibst ihn an eine URL in dem Format.
wo er sich die Keys Ed Runtime holen kann, oder du gibst ihm halt das Shared Secret per se mit. Das heißt, du kannst dich entscheiden, ob du den JVT validiert haben willst gegen einen IDP für ORF Sachen oder ob das sozusagen mit einem Secret. Ich würde halt den Weg.

  54:21
Yeah.
Mhm.
Mhm.

   54:38
preferen, weil der halt der Secure ist. Weil du halt sozusagen ein eigenes JVT-Issuring-System hast, was sozusagen auch Revocation und so weiter handled. Aber per se könntest du ihm auch einfach nur als JVT hier einen Shared Key geben oder du kannst

  54:48
Mhm.

   54:53
Könntest du auch auf 'Indication' auf 'Nun' drehen, wenn du sagst, dir ist alles wurscht?

  54:58
Ja, ja.
Könntest ja, ich hab den jetzt einmal.

   55:00
Ja.
Geschickt.

  55:08
Abgeschlossen hast du da jetzt irgendwas gekriegt.

   55:10
Ja, danke ***. Nein, aber hey, also da seh ich, kriegt er nicht immer. Oh, es ist schon Punkt 15:00 Uhr. Hast du noch Zeit oder bist du?

  55:19
Ich hab noch a bissl Zeit, ja.

   55:21
Wunderbar, dann kommen wir kurz zu den zu den Fragen zu den Metafragen, die ich bis jetzt noch nicht gestellt habe. Was ist deine Current Position in der *** oder generell was, was, wie ist dein Jobtitel?

  55:33
Der Fuchs, ganz normal, der Fuchs.

   55:34
Defop, wie lange machst du das schon? Defop sein, nicht bei der Firma, sondern generell.

  55:39
In.
Uh.

   55:46
War noch zu der Zeit, wo wir Mainframes hatten.

  55:49
nein, nein, nein, bitte, so weit bin ich jetzt nicht, nein, es ist, es ist schwierig, weil Defog als Begriff gibt es ja auch erst seit die, seit die Anfang die Zehner Jahren und vorher war es halt Operator oder in in Infrastructure. Mein Jobtitel ist es seit

   55:57
Ja, generell, ne.
Operations.

  56:07
war 1000
1415 offiziell und davor war ich halt auch im operativen Betrieb tätig. Hab dazu was Ähnliches gemacht.

   56:16
Versteh also, verstehe also, wie viele Jahre insgesamt?

  56:20
Von 2005 weg.
Grob Koston.

   56:24
Also.
20 Jahre

  56:27
Ja, gib mir Daumen.

   56:29
Passt, bist ein Experte, ***. Dafür sag ich danke. Hast du tägliche Exposure zur Application Configuration zur Application Configuration und Secret Management?

  56:41
Ja, würde ich schon sagen, oder?

   56:42
Yeah.
Bist du mit dem aktuellen Approach zufrieden, wie Secrets gehandelt werden?

  56:49
Wie wir die Secret Panel.

   56:52
Per se, ja wie gesagt, ja, wie der aktuelle Approach, mit wie du mit Secrets umgehst. Ich will da keine Erklärung, wie ihr das macht, weil das ist ja hochkompetential und das wollen wir nicht, aber ob du einfach damit zufrieden bist.

  56:52
Oder was meinst du?
Ja.
Mhm, ach so geht es ausreichend und ausbaufähig.

   57:12
Passt. Gut, wenn du hast vom X. L. Sikret Operator schon mal gehört, das haben wir eh schon established. Du findest das Mental Model auch relativ interessant. Genau, was also was siehst du sozusagen als Hauptstärke

  57:20
Mhm.

   57:28
This for some secret retriever tools with an externen secret operator.

  57:37
im Endeffekt, dass du das Secret nur da ablegen musst, wo sie wirklich benötigt wird. Sprich, es purzelt auf keinem Entwickler P. C. im Atom, auf keinen Default P. C. und im Optimalfall nur im Secret Store, dort wo es drinnen ist.

   57:48
Mhm.

  57:51
Und auf dem Server dort, wo es verwendet wird.

   57:51
No.
Verstehe, wunderbar. Dann waren das die paar organisatorischen Fragen, die ich halt brauche. Gut, genau. Komm jetzt zu den Open Questions. Ich lese sie dir jetzt vor oder lese sie dir selber durch, wie du willst. Und wenn es eine gibt, zu der du gerne etwas sagen würdest, dann nehme ich da sehr gerne deinen.
Deine Insights mit in meine Evaluierung.

  58:15
Das heißt, wenn ich zu diesen Fragen etwas sagen würde, wie.

   58:17
Yeah, genau.
Noch ein zweites Leitet Fragen. Alles gut. gibt ganz viele Fragen. Ich mag Fragen.

  58:27
Es ist so spät.

   58:30
Mhm.
Huhu.

  58:38
Man, die Concerns, die habe ich eh schon ein bisschen.
Ports auf einen Linux-Server. Direkte Exposion ist immer ein bisschen schwierig, weil das natürlich auch Angriffsfläche bietet.

   58:48
Ja.

  58:54
ist es dann als, ich denk jetzt noch 'nen Schritt weiter, wenn du jetzt sagst, O. K., das willst wirklich irgendwo Betrieb betreiben, ist es dann Open Source, wird das gewartet, gepatcht et cetera oder oder wie.

   58:57
Huh.
Jo.
Per se zur Zeit ist es ein Working Prototype. Wenn du ihn einsetzen willst, musst du also per se die Sources sind Open für das Ding. Heißt, du musst es entweder selber maintainen.

  59:11
Genau.
Aha.

   59:18
Oder aber sozusagen, wenn wirklich der Use Case kommt, dass man den verwenden will, kann man mit mir reden.

  59:26
Genau, aber das wäre für mich zumindest irgendetwas, wo ich sage, das muss gewartet werden, das muss gepatcht werden, muss Sicherheitsupdates geben, et cetera, et cetera, genau.

   59:28
Ja, ja.
Yeah.
Ja, hundertprozentig. Okay, passt. Wunderbarer Punkt.

  59:52
How acceptable is the security tradeoff of writing secrets as plain text files owned by the agent user?
Was meinst du damit

   1:00:02
Na, wie sehr würdest du sagen, dass es ein Acceptable Security Tradeoff, dass es eine Datei irgendwo am Host gibt, wo die Sachen drin stehen im Plaintext? Weißt du, mein, das ist ein O. K. A. Tradeoff, weil du dadurch sozusagen einfach Files hast und Files halt super einfach zum Einbinden sind von allen Applikationen, würdest du sagen, ist zwar 'n schönes Konzept, aber das ist viel zu
Insecure.

  1:00:23
Weil die Files, weil weil das XML dann trotzdem neues Plaintext für die Applikation.

   1:00:26
Genau, weil wenn jemand an die Daten der VM rankommt, sind die Daten da.

  1:00:35
Jo.
Ist jetzt meines Erachtens weniger was, was jetzt dein Prozess betrifft, sondern eher applikationsspezifisch betrifft.

   1:00:48
Ja, demnach kannst du also du würdest meinen, es ist ein Trade, auf dem man machen kann, wenn es die Organisational Constraints wollen.

  1:00:49
Du.

   1:00:57
Oder brauchen.

  1:00:58
Genau, man muss halt immer den Ist-Stand anschauen und dann tut, wo man hin will. Wenn der besser ist, dann kann man sagen: "Ja, das passt." Wenn er nicht besser ist, dann hat man vielleicht schon eine bessere Lösung.

   1:01:02
Ja.
Steht.

  1:01:21
Es gibt jetzt dann speziell für für wirklich Azure Cloud native Lösungen, gab es dann vermutlich schon noch bessere Lösungen, wenn man dran denkt, man hat dann managed Identity und kann das wirklich Cloud-seitig lösen, aber das das ist glaube ich bei dir auch schon wieder ein bisschen out of scope oder?

   1:01:31
Genau, aber man.
Die Sache ist, Manage identities könnte man rein typisch in den Agent für die Authentications and Secret Stores noch immer einbauen.

  1:01:43
Mhm.

   1:01:44
Rein tätisch könnte man das ja dann so machen, dass man es einfach so konfiguriert, insofern das am Linux-Host verfügbar wäre. Und pro Applikation, du könntest halt rein tätisch einen Store pro Applikation konfigurieren, der sich halt dann wirklich nur um die Datei kümmert. Aber das ist sozusagen nicht enforced, aber das du könntest ihn in die Richtung weiterentwickeln.

  1:01:51
Mhm.
Hm.
Also, ich sehe das ganze Konzept von dir jetzt so ein bisschen wirklich speziell für Applikationen, die das als Plaindex braucht und wie kann man das so, so sehe ich sie jetzt.

   1:02:05
Ja.
Genau.
Genau, genau, es ist sozusagen, ist es ein Trade-off und schön ist halt, dazu kommen wir dann auch gleich. Ja, genau, gut, dann sozusagen, nachdem die Zeit uns ein bisschen davonrennt. Da gibt es über Secret Management irgendetwas, wo du diesen Agent auch einsetzen.
Würdest, wenn also vom vom Konzept her.

  1:02:38
Über Secret Management. Du meinst, wenn ich sage, ich will den Agent verwenden, aber kann Secret damit verwalten, also ausrollen?

   1:02:44
Generell, ob du eine Idee hast oder wo man dieses Ding noch.
noch einsetzen könnte.
Wenn dir nichts einfällt, dann fällt dir nichts dazu ein.

  1:02:59
Na ja, configuration management generell overdraw, weiß nicht.

   1:03:01
Ja, verstehe. Ja, ja, hätten die anderen auch gesagt. Ja, du könntest ihn halt auch rein theoretisch für für für Linux Systemdateien Wartung einsetzen.
Also könnte es sein, dass alle großartigen Files, die du halt auf einem System brauchst, mit dem Ding abbilden, genau ja.

  1:03:24
Genau richtig, du musst die Templates generieren und dann der Rest ist im Prinzip komplett gleich.

   1:03:27
Genau, genau. Und halt, dass a production rate ist, braucht halt bissl Erfahrung und halt Maintenance und Adoption Rate und so weiter. Genau. Das the biggest Benefit, das haben wir eh schon besprochen, Limitations und Weaknesses natürlich.

  1:03:36
Yeah.

   1:03:44
Das alles lokal herumliegt und das Bootstrap Secret.
Das hast du eh schon beantwortet.
Ja, das die Frage wär vielleicht noch interessant.

  1:03:59
A well does that declarative kids based approach a line with your this verstehen it?

   1:04:06
Wie gut würdest du meinen, passt dieser deklarative Approach von dem Workflow mit Gitbased Sachen und C.I.C.D. Sachen?

  1:04:11
Ach so, mhm.
Was meinst du mit Gitbase Approach?

   1:04:18
Dass du sozusagen diese Secret Configuration in den Git Repository eintrickst.
Also das sozusagen.

  1:04:28
Komm, mach, steht ja nichts drinnen, oder?

   1:04:31
Ja.

  1:04:31
Das war ja das, was ich anfangs gesagt habe. Mit dem kann gewartet werden von.

   1:04:33
Genau, genau das, der der, wie du den Agen per se verwendest, ist deine Sache, aber ja, per se ist ja mein Workflow der.

  1:04:38
Yeah.

   1:04:42
Das ist, dass ja die Idee von dem, was ich propose, ist halt dieser Workflow, dass das halt super cool ist, aber das fehlt, demnach bau ich das und evaluiere das den Teil von der Equation.

  1:04:51
Mhm, mhm.
Genau, weil du musst ja in der in der in der OpenShift musst ja auch die ganzen Secret Server irgendwie also die.

   1:04:58
Yeah.

  1:05:00
Secret Konfiguration weiß jetzt nicht, wie es heißt.

   1:05:02
Ja.

  1:05:03
Warten und ausrollen und beim Linux-Approach funktioniert es gleich, würde ich behaupten.

   1:05:07
Genau, wunderbar.
Genau, gibt es irgendwelche Features, die du sehen würdest, die noch hinzugefügt werden müssen, bevor wir das Ding probieren können?
Und was und gibt es sozusagen irgendeinen technischen Grund, was dich davon hindern würde, diesen Workflow zu machen?

  1:05:40
technischen Grund vielleicht nicht, aber es gibt genug organisatorische Gründe.

   1:05:43
Yeah.
Organisatorisch is is natìrlich.

  1:05:47
Du, du, du darfst nicht vergessen, wir sind im *** und da ist alles. Das sind solche Themen sehr schwierig.

   1:05:52
Ja, ja, du, es geht mir nicht immer nur immer nur darum, was du jobmäßig dazu sagst, sondern ja.

  1:05:57
Ja, ja, nein, nein, also.
Genau das Einzige, was wie gesagt, ist mit der CICD und Konfiguration über.
Rest ausrollen gefällt mir ehrlich gesagt nicht so schlecht.
Du warst heute also solche, also in unserem Fall solche Secrets über A. C. I. C. D. ausrollen, würde ich bisschen problematisch finden. Da braucht man vielleicht ein anderes System.

   1:06:30
Du meinst, du meinst die Store Configuration.

  1:06:33
Ja genau, das ist so ein bisschen, wo ich, wo ich mir denke, puh, das, das, das will ich nicht über den Jenkins machen, das will ich eigentlich auch nicht über GitLab machen. Das könnte ich mir vorstellen, dass ich es lokal ausführe, nur dann ist die Authentifizierung vielleicht ein bisschen zu zu statisch oder so, dass man sie da.

   1:06:39
No.
Mhm.

  1:06:50
Userbased Authentifizierung oder so, dass dass dass ich sag, ich hab meine DevOps, jeder hat sein eigenes Secret. Wenn dem sein Secret liegt, dann tu ich halt nur das sperren und updaten. Da halt ein bisschen was in die Richtung, aber das ist das ist hängt ganz stark davon ab, wie man es.

   1:06:57
No.

  1:07:05
Wie man es einsetzt? Andere Firmen haben vielleicht CICDS, wo sie sagen da, da kann man Secrets drüber laufen lassen, dann dann passt das also wieder wie du es gemacht hast.

   1:07:07
Yeah.
Ja, ja, rein theoretisch, wenn also GitLab bietet dir auch Secrets, also rein theoretisch könntest du die Secrets, die dann applied werden, auch da abbilden und du könntest ja, wenn also voll durchautomatisiert könntest du ja sogar
Irgendwo einen Job haben, der die ganzen App Registrations per Terraform erstellt und die dann halt alle applied zum Agent, ohne dass die irgendjemand gesehen hat. Aber das ist sozusagen so Vendor und so Secret Store Dependent.
Wie diese Werte hinkommen, dass es sozusagen nicht wirklich den Sinn hat, das zu standardisieren.

  1:07:44
Mhm. Mhm.
Ja, das, das, das ist das Einzige. Wie gesagt, für unseren Case glaube kann man schon so so verwenden. Es ist halt mit Banksystemen halt praktisch schwierig.

   1:07:52
Verstehe.

 Transkription beendet
