
Transkription gestartet

  0:03
euch beiden möchte, beginnen oder ist euch egal?
Eine Person hört die Fragen halt zum ersten Mal, die andere hat währenddessen Zeit zum Nachdenken.

0:14
Ja, ist okay. Ja, passt.

   0:15
***.

  0:18
Also, hallo ***, was ist deine aktuelle Position, also jobtitelmäßig?

0:24
Falsch, o.
Top Titel ist schwierig, aber ich bin OpenShift Admin.

  0:30
Versteh.

0:30
Die Jobrolle ist System Engineer, Systems Engineer.

  0:34
Wunderbar, wie ja, wie lang machst du das schon?

0:35
Aber.
Ja.
Sex ja.

  0:43
Wunderbar.
Genau, musst du also, bist du täglich zu Secret Management Schrägstrich Application Management Exposed. Also, arbeitest du damit auf einer täglichen Basis oder ja, nein?
Dann.
Nein, Ton weg.
Und scheinbar ist *** weg.
So hello.

1:28
So tschuldigung.

  1:30
Alles gut.

1:34
O. K., irgendwas war mit musst du täglich.

  1:36
Ja, bist du täglich mit Application Schrägstrich Secret Management von Systemen Applikationen umgeben?

1:44
Und geben ja.
Verwenden. Es läuft ja im Hintergrund, also aber umgeben, ja.

  1:50
Versteh, bist du mit dem aktuellen Approach zufrieden?

1:57
Für die Plattform, ja, im Applications-Umfeld gibt es noch Verbesserungsbedarf.

  2:05
Verstehe, hast du schon mal mit External Secret Operator gearbeitet und wenn ja, was sind deine Meinungen zu dem Tool?

2:12
Ja, wir haben den im Einsatz.
Um.
Eigentlich.
Das ist eine elegante Lösung.
Weil es ja das Secrets Management prinzipiell.
im
Ganzen
Ja, Git Umfeld, wenn man es im Git da irgendwie ein Problem ist und ich find da external secrets operator oder external secrets.
Yeah, operate there.
Bietet eine elegante Lösung, um das heute.
Auch zu automatisieren, beziehungsweise.
Um diesen ganzen Headeck, den man da halt potenziell mitschleppt, ein Stück weit wegzukriegen.

  2:57
Mhm.
Wunderbar, gut. Dann haben wir da eine sehr ähnliche Meinung. Gut, dann die gleichen Fragen noch mal für dich, ***. Was ist deine aktuelle Position? Kriegst nicht Chop beschrieben.

   3:11
OpenShift Admin Systems Engineer.

  3:14
Wunderbar, ich würde euch in meiner Master-Phasis als Plattform-Engineers erwähnen, also.

   3:20
Ist OK.

  3:22
Ja genau, wie lange machst du das schon?

3:23
Ja, das passt.

   3:25
fünfeinhalb Jahre circa

  3:27
Passt. Hast du bei den anderen Fragen, also bei den Fragen, irgendwas anderes als der Simon zu Adden oder seid ihr da relativ gleich?

   3:37
Ich glaube, wir sind da sehr ident. Wir sind sehr happy mit dem external secrets Operator jetzt, weil er einfach Secret Management Gitups ready macht.

  3:48
Wunderbar, gut. Dann kommen wir mal dazu. Gut, ich brauch euch nicht erklären, was ein Central Secret Store ist. Die Info habt ihr. Ihr versteht, was Textual Secret Operator macht und wie er konfigurierbar ist. Das hilft mir extrem. Das nimmt uns allen Zeit raus aus dieser Equation.
Gut, das habe ich ja damals alles bei uns mitbekommen und unsere Applikationen sind ja noch nicht alle Cloud Raid. Und demnach habe ich mich masterarbeitsmäßig auf die Suche gegeben, auf die Suche begeben nach einem Tool, was das Gleiche wieder External Secret Operator anbietet für Linux Hosts.
Gibt aber leider nichts. Es gibt sozusagen 3 gröbere Kategorien in diesen fertigen Systemen. Das eine ist sind sozusagen entweder die Applikation, muss ich die Secret selber abholen oder du hast irgendeinen Agent laufen, den du über eine lokale API aufrufen musst.
Also, du musst irgendeine Form von Klugcode schreiben. Das andere ist natürlich SSH-basierend, die Sachen einfach irgendwo hinlegen. Das geht halt gut, solang du Glück hast und halt nur 3 Umgebungen hast. Und der andere Weg ist halt sozusagen irgendwelche Custom Skripts zu zu verwenden, die halt sozusagen nicht wirklich monitorbar sind und auch nicht Audible.
sind. Genau, sonst gibt's halt noch den Vault Agent auch noch für Linuxsysteme, aber der hat halt das Problem, dass er nur den Haschikow Vault unterstützt und nur einen Vault pro Konfigurierung leicht unterstützt und halt auch nicht in dieses Sequence Store und.
Secret Template System aufgeteilt ist. Demnach habe ich dann in meiner Master Physis sozusagen einen Prototypen für so einen Secret Retriever gebaut. Genau, und habe sozusagen dadurch dann einen hybriden Workflow so definiert, dass es halt ein Git Repository gibt, wo man die Store Configurations und die Secret Configurations einchecken
kann und dann diese ICT Pipeline Applied, die halt entweder auf den Connective Cluster oder auf den Unmenged Linux Host und würde sozusagen dadurch die konzeptionellen Ideen vom XL Sequid Operator auf ein Linux System übertragen. Und wie das genau ausschaut, sehen wir hier.
Das heißt sozusagen, ich habe ein kleines Binary gebaut. Das hat eine SQLite Database Pro Host, das ist die gebundled kommt. Das heißt, das System braucht nicht einmal SQLite, das ist im Go-Binary drinnen. Es hat eine API, wo ich ihn konfigurieren kann und Secrets und Secrets Stores.
hinschmeißen kann. Und ja, genau das, was macht der da macht, so wie der Xcel Secret Operator, hat der eine Time Pro Leaf Pro Secret, dann dann rolliert es dir. Er hat den Atomic File Writer, weil es da wichtig ist auf Linux Systemen, dass die Dateien halt richtig gesoppt werden und nur dann gesoppt werden, wenn es auch
Funktionierende Dateien sind und es gibt einen Graphana / Promepheus Alert Manager Endpoint, wo man sozusagen die Daten seiner Agents scrapen kann, weil es halt doch ein Distributed System ist, was sozusagen ja dann desto mehr Hosts hast, desto schwieriger wird es dann zu managen und zu erkennen.
Genau. Dann ja, Secret Store, was ist denn in dem Secret Store drinnen? Store wie beim E. O. S. ist, du hast halt einen Namen, also 'n Typ, der unterstützt wird. Dann hast du die Connection Details, die auf Indication und beim Erstellen schaut, ob er sich halt hin verbinden kann, ob das per se mal jede credential sind, die du grad ablegst.
Genau.

   7:31
Wat?
Hallo.
Hallo.

7:40
Also.
Ja, nix machen.
Es wirkt so, als wär.

   8:06
Max.

8:07
Headset gestorben.

   8:10
Hm.

8:54
Hallo.

   8:55
No.

8:56
***, du hörst Bischof, gell?

   8:57
Ich höre dich.

8:58
Ja, passt.
K, da geht's einfach weiter.

   9:03
Ich glaube, ich glaube, ich glaube, er ist gerade an meine Hustle.

9:25
Schaut aber sehr durchdacht aus.

   9:27
Mhm.
Mhm, less is my body outsickay. Vielleicht make this done.

9:54
Ja.
Hallo.
Hello.

   10:23
Ja.
Äh, lassen wir mal sicher, ich ruf mal extra Nummer.

10:36
Ah, passt.
Ich denke, du hast nicht erreicht, oder?

   11:32
Na, es ist gleich gefordert worden.

11:36
Ja.

   11:37
Probier mal mal, wie lange es so ist.

11:42
Lang, lang, kurz, kurz, lang, lang, was?

   11:43
No.

11:58
Hier her, an dich nicht.

   11:58
Wir hören dich nicht.

  12:06
Bono.

   12:07
So, jetzt hören wir die wieder.

  12:09
Ono, Ono, wo seid ihr ausgestiegen? Oh Gott!

   12:11
Ja, das ist du hast jetzt glaub ich, du hast jetzt glaub ich 5 oder 6 Minuten a lagered.

  12:16
Oh Gott, Gott, ja, wo war der letzte Punkt?

12:17
Yeah.
Da das haben wir, das haben wir noch mitgekriegt, das haben wir noch mitgekriegt.

   12:19
Uh, da waren wir genau. Da waren wir. Uh, matrix, matrix endpoints, sust the available is ferrometers.

  12:21
Okay, wunderbar.

12:24
Yeah.

  12:24
Ja, oh Gott, das, das tut mir urleid. Ich hab es Team, ich war einfach am Reden und Teams hatten wir, ah, jetzt muss ich noch einstellen, dass ich die Notification bekomme. Auf jeden Fall. Das, also das Konzept des Ages habt ihr einmal genau. Es gibt einen Store, es gibt Managed Files,

12:26
Yeah.

   12:27
Nein, kein Problem.
Äh.

  12:41
Das ist so wie Berlin.

   12:41
Genau das Managed File, das haben wir noch nicht. Das da da bist du dann weg gewesen.

  12:44
Okay, gut. Das Managed File ist sozusagen so wie das Exterial Secret auf Linux.

12:45
Yeah.

  12:50
Das ist halt sozusagen, kann ich halt angeben, wo es hin soll, hab sozusagen Store References und kann sozusagen so wie im Next Tile Secret Operator mit den Stores arbeiten. Dann hab ich die Rotation Policy, die halt sagt, wann wie roll ich das File und ich hab halt ein Output Format Jason Yummel 11 Pem.

12:51
Mhm.

   12:56
Mhm.

  13:08
und sozusagen, die werden halt überprüft, um halt sozusagen, wenn jemand etwas falsch escaped oder eine Datei nicht valide ist, dass die halt sozusagen verwendet werden kann. Genau, gut, die dieses File, also genau, der Agent PC hat eine Crowd API, mit der man alles steuern kann.

   13:08
Mhm.

  13:25
Ein Get Endpoint auch, aber da kriegst du niemals die Secret Values zurück, sondern nur die Konfiguration für Secret oder halt die vom Store. Aber mit den retected Infos für den Login, weil sonst wäre das ziemlich blöd, wenn wir die Login-Daten zurückgeben würden. Dann haben wir einen Audit Log.
Der geführt wird für jede Aktion. Dann haben wir die TTL, wie gesagt, den Metrics Endpoint und die Pfade, wo man hinschreibt, müssen explizit erlaubt sein. Man könnte natürlich sagen, der System User, den es gibt für den Agent, sollte da, sollte nicht mehr berechtigt sein, ja, aber vielleicht.
Kann der mehr oder hat oder sozusagen, du verwendest vielleicht für den Agent den gleichen System User wie den Application User. Ist zwar nicht vorgesehen, aber könnte man ja. Demnach kann ich es dann noch mal explizit unter Laun.
Genau so, ihr seid noch da.

   14:15
Mhm.

14:15
Yeah.

  14:16
Wunderbar, gut. Der Client, also der Agent sollte rein theoretisch auch ins Internet, aus dem Internet erreichbar sein können. Demnach habe ich mich für einen Layout of Indication Approach entschieden, dass es einen M.T.L.S. Track gibt und nach dem M.T.L.S. Track noch einen JVT Track gibt und der ist sozusagen entweder mit dem Shared Secret.
oder mit Identity Provider wie Kicklog oder einem anderen ORF Provider verwindbar, sozusagen, dass du eine doppelte Sicherheit hast, um das Ganze halt sozusagen Internet Facing Exposure zu erlauben. Dann haben wir noch den Atomic Fileswap.
Der Atomic Fileshop, den müsst ihr jetzt nicht genau verstehen, aber es geht halt einfach darum, er checkt, ob er überhaupt in dem Pfad sein darf und löst auch einen relativ konfigurierten Pfad auf, ob er da schreiben darf. Dann schaut er sich das evaluierte File an, schaut, ob das sozusagen zu den definierten Standards passt.
Dann macht er halt ein Temp-File, schreibt das Ganze mal, setzt die Permission, schaut, dass da alles passt. Und nur wenn all das funktioniert hat und das ablegen konnte und das File geschlossen werden konnte, geht er hin und Atomic Less Wap, das mit dem können wir halt generieren. Also kann ich garantieren.
Dass sozusagen nur Dateien ersetzt werden, wenn es auch wirklich funktioniert.
Genau dann, die Installation vom Agent, hab ich mir gedacht, löse ich über Ensible aus dem einfachen Grund. In Kubernetes haben wir unsere schöne Cube CTL, wo wir sozusagen deklarativ Sachen gegen das System werfen können und das Ding macht das einfach. Das fehlt uns natürlich auf Linux und ich hab gedacht.
Das ist sozusagen Ansible, da der beste Counterpart ist. Was braucht dieser Agent? Was macht das Playbook? Es würde dir ein System User kreieren. Es macht dir die Gruppe für alle User, die sozusagen die Sequels lesen dürfen. Dann schreibt es, dann macht es die Pfade, die du brauchst.
Template SD Configuration und das System D-Service und es lädt halt das Binary runter mit der Version, die du möchtest. Genau, dann die Konfiguration des Agents ist relativ einfach. Es gibt einen Note Name, es gibt die Trusted Roots, wo man halt definiert, wo er überhaupt hinschreiben darf.
über also sozusagen application-level-mäßig, nicht system-user-mäßig, dann die Login Config und dann haben wir noch die Security Settings. Ich kann TLS auf und abdrehen, wie gesagt, diese 3, das ist das Cert File, das Key File und das Client CA File sind sozusagen die
drei Sachen, die nicht Bootstrapper sind, die sozusagen schon am Host da sein müssen/selber schon da sind. Das ist sozusagen Outer von Scope from Research. Da gibt's halt die Authentication Methoden non, MTLS, JBT und Spire. Genau.
Man könnte natürlich nur MTLS auch für die Authentication verwenden. Was ich empfehlen würde, halt für Internet Facing Exposure, wär du hättest MTLS einfach nur mal zu checken, ob der Request valide ist und dann JVT, um sozusagen dann noch zu schauen, ob der JVT passt, genau.
Habt ihr schon mal mit Speyer gearbeitet oder wisst ihr, was das ist?

   17:07
Umm.
Wir habens uns einmal angeschaut, also nur nur als eine Demo von Red Hit bekommen, weil wir das auch irgendwann implementieren wollen.

  17:11
K, ja.
Verstehe, da habe ich mich auch ein bisschen herumgespielt mit Spyer von Spyer, um sozusagen rein theoretisch könnte man das Ganze dann auch Workload Identity Based in Zukunft umsetzen. Also die Middleware habe ich in Go damals geschrieben und herumgebastelt, aber mehr ist nicht.
Da genau, das ist sozusagen auch noch möglich. Gut, kommen wir jetzt zum Configuration Example. Vom Exile Secret Operator kennt ihr es eh gut, da brauch ich nicht zu viel erklären. Ich hab das Ganze so abgebildet, dass bei mir die ganzen Sachen halt anstatt hier über das halt über Typ gehen.
URL und das der ganze auf Indication Block ist dann sozusagen so gelöst. Das heißt, man kann diesen Teil nicht wirklich gut in Git einchecken, weil du sozusagen da die Secrets drin stehen hast. Das heißt, die müssen sozusagen in einem eigenen Secret Repository noch gemanagt werden.
War die Sache ist, es gibt halt keinen schönen und einfachen Weg, wie man das so abbilden kann, weil dadurch, dass du halt auf einem halt normalen Linux Host halt nicht diese Secret A. P. I. hast, die du halt beim Excel Secret Operator hast, kannst du schwierig auf etwas verweisen.
Was nicht wirklich da ist, genau. Und jetzt hier configuration mäßig, hier zum Beispiel eine kleine Configuration für den Secret Store. Also für ein kleines XML Secret wäre das hier für Excel Secret Operator. Das kennt ihr, glaube ich, relativ gut.
Ich hab sozusagen das Ganze ein bisschen compressed, ein bisschen umgenannt in meiner Implementierung, aber es passiert auf dem Gleichen auf dem auf dem Storenamen. Dann hat das Ding auch einen Namen, um halt sozusagen schnellere Lookups zu haben.
Dann gibt es einen Filepath.
Wo halt die Datei landen soll, das ist ja sozusagen bei Kubernetes auch nicht wichtig. Dann den Content und die Secret References. Was ich sozusagen beim Eugency Secret Operator immer sehr störend finde, ist dass du sozusagen so viel Metadaten brauchst für so ein einfaches Mapping. Das heißt, bei mir ist das Mapping, dass sie einfach das, was da steht, muss da stehen.
Und so muss es im Secret Store heißen.

   19:16
Mhm.

  19:17
Genau so viel zur Konfiguration. Nachdem ihr beide schon Excel Secrets generiert habt und geschrieben habt, glaube ich, dass ihr in der Lage seid.
Diese Sequels zu konfigurieren, dieses Slide ist viel mehr für die Leute, die sozusagen noch nie was im EOS zu nein ESO zu tun hatten, habe ich auch falsch geschrieben da um halt denen sozusagen ein bisschen ein Gefühl zu geben, aber ich glaube, das habt ihr beide.

   19:41
Mhm.

  19:43
Wunderbar, ja, dann kommen wir jetzt auch schon zu einem kleinen Form. Wenn euch irgendeine Frage unklar ist, einfach mit mir reden.
Sonst sind danach nur noch die offenen Fragen und dann sag ich auch schon vielen Dank für eure Zeit.
Soweit alles klar bei euch?

   23:44
Mhm.

23:45
Ja, ich bin gerade fertig.

  23:46
Drück mal.

   23:46
Turn off a two one gold.

  23:48
Wunderbar. Ihr seid die ersten, die mir keine einzige Frage zu den Fragen gestellt hat, gestellt habt. Jetzt ist die Frage.
Das ist sehr interessant, aber freut mich.

   23:58
Ja, so zu den Open Questions.

  24:00
Nein, nein, generell einfach zu den Fragen wurden mir auch schon oft Fragen gestellt.

   24:01
Not it.
Ach so, na ja, wir arbeiten mit dem, ja.

  24:08
Ja, wunderbar, wunderbar.

   24:08
Fast täglich.

24:09
Ja, ich hab gestern auch wieder, glaub ich, 2 verstellt, also.

   24:13
Ach so.

24:16
Ja.

  24:17
Gut, ich kann euch jetzt die Fragen natürlich vorlesen oder ihr sucht euch halt welche aus, die euch sozusagen gefallen, wo ihr dazu was hinzufügen wollt, wie es euch lieber ist.

24:41
Okay, ähm.
Ah.

  24:58
Sonst kann ich euch einfach auch die Fragen stellen, die mich am meisten interessieren aufgrund eurer Rolle.

   25:04
Local Summer.

25:05
No.

  25:06
Wunderbar, was also, wie seht ihr den Security Tradeoff als das Secret als Plaintextteil am Host zu haben? Schrägstrich die Condentials zum Secret Store in der SQLite Database drinnen zu haben.
Also, die Secrets zum Store.

25:27
Prinzipiell, also ich könnte das aus Linux Admin Sicht sprechen, nur wie das halt umlegbar ist in OpenShift oder so, oder halt wie es, wie es vergleichbar ist. Sehe jetzt gar nicht so dramatisch,
weil wenn man jetzt zum Beispiel den den Secret Store hernimmt, äh, den externen Secrets Operator und C. A. Secret sozusagen erstellen lässt, dann würde das ja ohne Maßnahmen im Kubernetes Etsy
Die ja unverschlüsselt in der im Netz liegen, sprich das ist also.
Wenn man nicht hergeht und sein SD, also die Datenbank in in Kubernetes verschlüsselt, ähm, liegt das Secret A im plain Text dort nochmal drum.

  26:09
Yeah.

26:17
In den Playtex, aber halt Base 64 inkludiert, aber das ist ja nichts. Also von dem her ist es.

  26:19
Ja.

26:24
Vergleichbar, also.

   26:26
Ja, solang die die File Permissions so passen, dass halt wirklich nur der der Agent oder halt system Users dann drauf zurückgreifen können und die das die Secrets wirklich benutzen müssen.

  26:42
Verstehe.

26:42
Yeah.

  26:44
Wunderbar, ja, das hier zurück.

26:44
die einzige, ja, die einzige Thematik ist halt, aber ich weiß nicht, ob das irgendwer verwendet in irgendeinem Form Linux wäre halt, wenn es das in einem Keyring auf Linux halt irgendwie reinbringst, aber ich glaub nicht, weil da brauchst
mit Debus Anbindungen und dann brauchst du schon wieder eigentlich ein U.I., damit das gescheit funktioniert. Das hast du auf die meisten Server nicht. Und dann ja, also ich, ich würde den Trade auf wahrscheinlich
In Kauf nehmen an die, die meistens sind, also die meisten.
Sachen ist halt.
Zu die anderen Fragen, vielleicht meistens hat man ja schon irgendwas und dann ist halt schwierig zu argumentieren, warum man jetzt, warum das Neue jetzt so viel besser ist als wie das Euge. Wenn ich jetzt zum Beispiel hergehe und sag, ich hab meine Ensible, meine Passwörter als Ensible Valls im Ensible.
Und es funktioniert. Warum sollte ich dann wechseln? Aber so für so, sagen wir mal, wenn man etwas Neues macht, ist das wahrscheinlich die bessere Lösung.

  27:53
Ja, das Argument wäre passiert, dass du halt diese kurze Time-to-Lief abbilden kannst. Das heißt, du könntest halt alle 5 Minuten ein neues Secret haben, ohne alle 5 Minuten das Ensible Playbook zu starten, was das Gleiche macht, aber ja, ja.

28:04
Das stimmt natürlich, ja.

   28:05
Mhm, also die Rotation ist auf jeden Fall ein großer Vorteil.

28:07
Ja, das stimmt.

  28:09
Und dass halt jeder es.

   28:09
Every hops, when I there that secret stores are eben so unpeated.

28:13
Mhm.

  28:14
Ja, und was halt auch relativ cool ist, warum ich den Approach sozusagen über so as ensible S. S. H. based Approaches sehe, ist der einfache Grund, dass du sozusagen auch Developer ihre Sequels managen lassen kannst, weil keine Ahnung, wie es was ihr da genau geantwortet habt, das soll ich eh später sehen.

28:27
Mhm.

  28:31
Ja, aber sozusagen, die du kannst ja das Repository mit den Secret Configs jedem geben, weil du sozusagen ja da nur Key Value Sachen hast und sozusagen ein Developer ist niemals mehr exposed zu den Secret Values, aber kann sich selber seine eigenen Konfigurationen managen.

   28:38
Mhm.

28:39
Mhm.

  28:46
ohne viel, da seh ich sozusagen den größten Nutzen in diesem Workflow mit, dass du halt

28:52
Wenn also ich, ich kann ja über das Tool wahrscheinlich auch, weil ja alles mögliche ausrollen kann, kann ich auch Zertifikate ausrollen, oder?

  28:59
Ja, du kannst auch PEM-Files konfigurieren.

29:01
Ja.
Weil da Siggi eigentlich.
ein Riesenvorteil, wenn man jetzt sieht, dass Zertifikate immer kürzer valide sind. Ruft ja, glaub ich, Nummer 7 Tag, 5 Tag, 6 Tag, keine Ahnung, ist irgendwas. Und du das jeden Tag halt austauschen oder halt von

  29:14
Ja.

29:23
oder du halt man halt selber ein eigenes gehostete P.K.I. hat und die Anforderung hat, diese Zertifikate
Oft zu rollieren.

  29:33
Yeah.

29:33
Dann ist es natürlich von Vorteil, wenn es einfach sozusagen.
Selber passiert, als wie dass man es immer antriggern muss, ja.

  29:38
Audit yeah.
Genau, und das ist halt Auditabelle ist und man sozusagen halt einen genauen, auch genauen Trace hat, wann welches Secret, wo, wie gerollt wurde.

29:53
Genau, ja.

  29:54
Genau gut, dann sozusagen gibt's ein Szenario, wo ich vor, wo ich sozusagen vorstellen könnte, jetzt den nicht zu verwenden.
Das heißt, jetzt aus eurer Sicht, um oder besser gesagt, andere interessantere Frage für mich: Könnt ihr euch vorstellen, dieses Tool zu verwenden, um auf Servern, die halt sozusagen dann ein OpenShift-Cluster abbilden?
Dort Dateien zu managen.
Weil Kubanetes ist braucht ihr eigentlich auch Bootstrap secrets, oder?

   30:26
No.

30:30
Ja, das ist im OpenShift, also Kubernetes per se, ja. Aber OpenShift braucht das nicht, weil die sozusagen direkt mit die Cloud Anbieter, sei es WIS, sei es Azure, direkt über die A.P.I. mit denen redet und sie sozusagen das.

   30:31
Das das Lauf komplett anders.

  30:34
Ah.

30:46
Selbst ausmacht.

  30:47
Verstehe.

30:47
Auf also für uns als Openshiftadmins wahrscheinlich nicht, weil wir haben ja den externe Secrets Operator und der Riesenvorteil von dem für uns ist, dass der vom Hersteller supported wird. Also das ist das, das ist halt das.

  30:52
Mhm.
Yeah.

31:03
größte Team und fìr uns ist halt wichtig, dass wir einen Ansprechpartner haben, der das Fìr uns halt tut, äh, Fìr Linux Use Cases.
Kann man sich das natürlich vorstellen, aber dann müssen wir also ich weiß nicht, ob du noch Gespräche mit den Kollegen aus dem Linux-Umfeld führst, wahrscheinlich die.

  31:21
Ich werde es dem *** heute noch präsentieren. Per se, ob es bei uns adoptet wird, sozusagen. Es ist.

31:24
Ja.

   31:25
Also.
Falls du, falls du Ansprechpartner haben willst, können wir da gern.
Das Weiterleiten also.

31:32
Also wirklich die, die die die Linux Admins Admins also der *** ist ja a.

  31:36
Ah, OK.

   31:37
No.
Wir haben ein eigenes komplettes Linux-Team, also die kannst du definitiv auch.

31:41
Yeah.

   31:47
Ansprechen, das war das Team vom ***, da zum Beispiel.

  31:48
Versteh.
OK.

   31:55
Der *** ist einer der.

  32:00
was werde ich

   32:00
Und da, da *** zum Beispiel, das soll übers Richtung unter ***, übers Richtung Anselbüll und und Rolte und so weiter und Dinge wahrscheinlich am meisten machen.

  32:11
Steht ja, werd ich schauen, weil per se glaub ich, hab ich schon die Zahl durch, die ich brauch für die Masterarbeit aber ja.

32:16
Ja, natürlich, ja.

   32:18
Ja, also die haben sich einfach gute Inputs genau für das da.

32:19
Ja, ja.

  32:22
Ja, ja, verstehe. Wunderbar, gut, dann.
Darf ich wieder was machen, danke. Genau. Hier sind noch 'n paar, ja sozusagen, seht ihr es also, was seht ihr sozusagen als als Hurdles von diesem Fundament Approach, dass wir das sozusagen beides, also.
Seid ihr ein Fan davon, sozusagen diesen hybriden Workflow so abzubilden mit einem Agent und dem Excel Secret Operator, oder meint ihr das sozusagen? Also, was ist generell eure Meinung dazu? Kürzen wir das einfach ab.

   32:58
Auf jeden Fall, weil im Endeffekt, du kannst ja dann hergehen und sagen, meine Secrets werden einfach von 2 Agents, egal wo der Agent lauft, ob da jetzt im im Cogoniz oder am Linux Host lauft.
wird immer das selbe Secret dann verwendet. Äh und das so solange das manchmal noch immer wieder automatisiert rotiert wird, hat das natürlich einen kleinen Konfigurations beziehungsweise Overhead-Vorteil, weil man sie einfach um weniger
Oder weil weniger Secrets auslaufen können, zum Beispiel, oder Zertifikate auslaufen können, und man muss sich auch um weniger kümmern.

33:34
Yep.
He is at.
Ja, ich seh den Riesenvorteil halt eigentlich, dass man sie auf an Keystore sozusagen fixieren kann. Man kann sagen, OK, wir verwenden jetzt Azure Keyword.
Und.
hab da alle meine Secrets drin. Ich bin in der OpenShift und hab aber auch meine Linux-Server und muss aber sozusagen mein Tooling für Secrets, sagen wir mal, austauschen oder was, eigentlich alles nur im Azure Keyword machen. Und
Alles passiert dann automatisch. Das Ding wird ausgefällt auf die über 'n externen Secrets Operator auf Openshift und der Linux Secrets Agent Heut macht das halt für die Linux Server und ihm als also.

  34:18
Agent.

34:26
Ich kann es mir vorstellen, wir sind keine Applikationsentwickler und wir haben keine Linux-Server jetzt in dem Sinn, aber da kann ich mir halt vorstellen, ich muss sozusagen nur schauen, dass im A.M. Keyword oder halt auch mehrere Keywords oder was auch immer ich da abbilden will, nur mir auf den A. P.

  34:28
Yeah.

34:43
Beschäftigen und muss mich nicht drum kìmmern, wie ich das jetzt groß auf alle verschiedenen Systeme ausrollen muss. Also, das ist schon Vorteil.

  34:49
You know.
Genau. Cool. Gut. Dann sozusagen, bevor ihr sozusagen euch vorstellen könntet, so ein Agent in Produktion einzusetzen, was müsste dieser Agent machen? Schrägstrich, welche Features würden euch jetzt sozusagen noch fehlen, die sagen, das wäre für euch per se ein Roadblocker, natürlich Testing und Testing und sozusagen.
Das beweisen, dass er funktioniert und dann halt klein ausrollen, das ist mir schon klar. Aber habt ihr irgendwas gefunden, was sozusagen davor noch dazu kommen müsste?

   35:21
Bescheid.

  35:32
Nein, meine Audio ist da, oder?

35:34
Ja, ja, es ist ja.

   35:34
Ja, du bist bist da hinten am Nachdenken.

  35:34
Ja, ja, wunderbar.
Alles gut.
Seit dem Feuer bin ich ein bisschen paranoidiert.

   35:41
That's good.
Also in Produktion ist für uns meistens sehr wichtig, eben dass es irgendwo supported ist. Das ist halt unser, unser Ansatz als als kleines Plattformteam. Wir versuchen, das alles so gut wie möglich supported ist. Wir haben natürlich auch Sachen, die

  35:51
Ja, ja.

   36:03
Wo wir sagen, ja, wir sind ja Open Source Tooling, das wir verwenden, ja, das schon, aber die zentralen Themen versuchen wir eben so gut wie möglich supported die Varianten zu haben in Produktion.

  36:12
Ja.

36:18
So.

   36:19
Genau, das ist also Thema auf jeden Fall.

36:23
Genau, was jetzt natürlich nicht heißt, dass jetzt ein Entwicklungsteam, weil das das und hergehen kann, also jetzt nicht bei uns, aber halt sagen wir mal diesen Secret Agent hernehmen kann und für sich einfach ein paar Secrets managen kann.

  36:35
Yeah.

36:38
Es ist halt nur bei uns so, wenn wir halt was ausrotten, dann ist es halt so zentral, dass und und wir so wenig Leid, das ja.

  36:42
Ja, du.
Generell, generell muss das nicht sein, ob ihr es per se in eurem Environment, sondern ob ihr sozusagen, also keine Ahnung, für ein Homelab oder für eine imaginäre Welt, ob es Features gibt per se, die ihr meint, das Ding würde, also das würde noch unbedingt brauchen, Schrägstrich.
Oder wir meinten, nein, das könnten wir so mal testen, so wie es konzipiert wurde.

37:03
Ach so, ja.
Also, da sehe ich dann eigentlich nichts.

   37:06
Ach so, wie es konzipiert ist, auf jeden Fall.

37:08
Yeah.

  37:09
Wunderbar, cool.

37:10
Also, hier gehe ich keine Thematiken.

  37:12
Cool, cool. Ja, kann ich auch schon berichten. Das erste Secret auf einem Linux-Server wird jetzt schon gemanagt, damit und die erste Applikation läuft damit.

37:21
Sehr good.

   37:21
Nice, voll cool.

  37:22
Also schauen wir mal, wo es in welche Richtung jetzt da weitergeht mit dem Tool. Halt noch alles sehr unoptimiert und mit Curl befehlen über meinen User auf dem Server, aber wie all Star Zangl.

   37:32
Sodom.
Dann hätte ich auf jeden Fall noch die letzte Frage, die auf jeder Konferenz kommt. Falls du das Ding mal vorstellst, ist es Open Source.

  37:37
Yeah.
Ja, es ist Open Source.

37:45
Ja.

  37:46
Und sogar die Dokumentation ist alles ist Open Source seit dem Moment. Also, I don't keep. Ich gebe alles zurück an die an GitHub. Alle E. R. Modelle dürfen lernen von mir alle.

   37:49
Ah, stark.

37:50
Yeah.
Yeah.

   38:00
Cool.

38:01
Ja, letztens waren wir eh wieder auf einer auf einer Konferenz und dann haben andere, also von andere Firmen, Leute was vorgestellt und alle haben dann immer gefragt, ist das Open Source und die dann immer? Na, also in dem Fall, das ist schon, also das ist schon viel wert.

  38:02
Genau, ziehen wir das kurz mal.
But.
Genau.
Das das Ensemble Skript mit den ist ist da generell der ganze Source Code und es gibt auch es gibt auch Documentation, die sicher nicht mit der I. mit generiert wurde, die natürlich alle händisch geschrieben, wie man das heutzutage macht, aber die ist auch über GitHub Pages released. Das heißt ja, das Ding ist komplett offen und

   38:17
Sehr cool.

38:28
Ja, klar.

   38:28
Mhm.

  38:32
Kann verwendet Extended und so weiter werden und ich werd wahrscheinlich bald auch in unser GitLab reinmiral.
Und sozusagen die Releases dann auch probieren darzubauen, genau.

   38:45
Hold boss.

38:46
Ja, cool, richtig klasse.

  38:47
Wunderbar, ich sage vielen Dank für eure Zeit. Es sind wieder wichtige Inputs und Daten, weil das ist das, was ich für die Masterarbeit noch brauche.
Ich stör.

Transkription beendet
